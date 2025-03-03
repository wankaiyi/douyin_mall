package release_real_quantity

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/product/biz/dal/redis"
	productModel "douyin_mall/product/biz/model"
	"douyin_mall/product/infra/kafka/constant"
	rpc "douyin_mall/product/infra/rpc"
	"douyin_mall/rpc/kitex_gen/order"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type ReleaseRealQuantityHandler struct {
}

func (ReleaseRealQuantityHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ReleaseRealQuantityHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ReleaseRealQuantityHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	ctx := session.Context()
	for msg := range claim.Messages() {
		klog.Infof("消费者接受到消息，Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s \n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		session.MarkMessage(msg, "start")
		value := msg.Value //消息内容
		var dataVo string = ""
		//将消息反序列化成Product结构体
		_ = sonic.Unmarshal(value, &dataVo)
		msgCtx := otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))
		_, span := otel.Tracer(constant.ServiceName).Start(msgCtx, constant.AddService)
		err := ReleaseRealQuantity(ctx, dataVo)
		if err != nil {
			klog.CtxErrorf(msgCtx, "消费者处理消息失败，err=%v", err)
			return err
		}

		span.End()
		session.MarkMessage(msg, "done")
		session.Commit()
	}
	session.Commit()
	return nil
}

func ReleaseRealQuantity(ctx context.Context, dataVo string) (err error) {
	//根据orderId查询订单信息
	orderData, err := rpc.OrderClient.GetOrder(ctx, &order.GetOrderReq{OrderId: dataVo})
	if err != nil {
		return err
	}
	//先判断订单是否被消费
	orderKey := "product:order:" + dataVo
	//十分钟过期
	ensureScript := `
		local k=KEYS[1]
		local a=redis.call('incr', KEYS[1])
		redis.call('EXPIRE ', k, 600)
		return a
	`
	result, err := redis.RedisClient.Eval(ctx, ensureScript, []string{orderKey}).Result()
	if err != nil || result == nil {
		klog.CtxErrorf(ctx, "判断订单有无消费时异常，err=%v", errors.WithStack(err))
		return err
	}
	klog.CtxInfof(ctx, "释放真实库存，订单id=%v, result=%v", dataVo, result)
	//只有为1的时候才能消费
	if result.(int64) != 1 {
		klog.CtxInfof(ctx, "订单已被消费，不再处理")
		return nil
	}
	//从订单信息获取商品信息列表，其中包括各个商品的id和购买的数量
	products := orderData.Order.Products
	luaScript := `
		local function process_keys(keys, quantities)
			for i, key in ipairs(keys) do
				local quantity = tonumber(quantities[i])
				local stock = tonumber(redis.call('HGET', key, 'stock') or 0)
				local lock_stock = tonumber(redis.call('HGET', key, 'lock_stock') or 0)
				if stock - lock_stock <= quantity then
					return 2
				end
			end
			for i, key in ipairs(keys) do
				local quantity = tonumber(quantities[i])
				redis.call('hincrby', key, 'stock', -quantity)
				redis.call('hincrby', key, 'lock_stock', -quantity)
			end
			return 1
		end
		return process_keys(KEYS, ARGV)
	`
	keys := make([]string, 0)
	args := make([]interface{}, 0)
	for _, pro := range products {
		keys = append(keys, productModel.StockKey(ctx, int64(pro.Id)))
		args = append(args, pro.Quantity)
	}
	//开启事务，批量扣减商品的真实库存和锁定库存
	result, err = redis.RedisClient.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		return err
	}
	if result.(int64) != 1 {
		klog.CtxErrorf(ctx, "库存不足")
		return errors.New("库存不足")
	}
	return nil
}
