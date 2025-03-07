package reduce_real_quantity

import (
	"context"
	common "douyin_mall/common/constant"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/product/biz/dal/redis"
	productModel "douyin_mall/product/biz/model"
	"douyin_mall/product/infra/kafka/constant"
	rpc "douyin_mall/product/infra/rpc"
	"douyin_mall/rpc/kitex_gen/order"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type ReduceRealQuantityHandler struct {
}

func (ReduceRealQuantityHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ReduceRealQuantityHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ReduceRealQuantityHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	ctx := session.Context()
	for msg := range claim.Messages() {
		klog.Infof("消费者接受到消息，Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s \n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		uuidStr := uuid.New().String()
		context.WithValue(ctx, common.TraceId, uuidStr)
		session.MarkMessage(msg, "start")
		value := msg.Value //消息内容
		var orderId string = ""
		//将消息反序列化成Product结构体
		_ = sonic.Unmarshal(value, &orderId)
		klog.CtxInfof(ctx, "orderId:%v", orderId)
		msgCtx := otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))
		_, span := otel.Tracer(constant.ServiceName).Start(msgCtx, constant.AddService)
		err := ReduceRealQuantity(ctx, orderId)
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

func ReduceRealQuantity(ctx context.Context, orderId string) (err error) {
	//根据orderId查询订单信息
	klog.CtxInfof(ctx, "开始释放真实库存,订单号是%v", orderId)
	orderData, err := rpc.OrderClient.GetOrder(ctx, &order.GetOrderReq{OrderId: orderId})
	if err != nil {
		klog.CtxErrorf(ctx, "rpc查询订单信息失败，err=%v", errors.WithStack(err))
		return err
	}
	//先判断订单是否被消费
	orderKey := "product:order:" + orderId
	//十分钟过期
	klog.CtxInfof(ctx, "开始判断订单是否被消费,订单号是%v", orderId)
	ensureScript := `
		local k=KEYS[1]
		local a=redis.call('incr', KEYS[1])
		redis.call('EXPIRE', k, 600)
		return a
	`
	result, err := redis.RedisClient.Eval(ctx, ensureScript, []string{orderKey}).Result()
	if err != nil || result == nil {
		klog.CtxErrorf(ctx, "判断订单有无消费时异常，err=%v", errors.WithStack(err))
		return err
	}
	klog.CtxInfof(ctx, "释放真实库存，订单id=%v, result=%v", orderId, result)
	//只有为1的时候才能消费
	if result.(int64) != 1 {
		klog.CtxInfof(ctx, "订单 %v 已被消费，不再处理", orderId)
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
				redis.call('EXPIRE', key, 300)
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
	klog.CtxInfof(ctx, "开始释放真实库存,订单号是%v, 商品id列表是%v, 数量列表是%v", orderId, keys, args)
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
