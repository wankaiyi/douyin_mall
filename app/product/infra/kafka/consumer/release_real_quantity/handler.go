package release_real_quantity

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/product/infra/kafka/constant"
	"douyin_mall/product/infra/kafka/model"
	rpc "douyin_mall/product/infra/rpc"
	"douyin_mall/rpc/kitex_gen/order"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
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
		dataVo := model.ReleaseRealQuantitySendToKafka{}
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

func ReleaseRealQuantity(ctx context.Context, dataVo model.ReleaseRealQuantitySendToKafka) (err error) {
	//根据orderId查询订单信息
	_, err = rpc.OrderClient.GetOrder(ctx, &order.GetOrderReq{OrderId: dataVo.OrderID})
	if err != nil {
		return err
	}
	//从订单信息获取商品信息列表，其中包括各个商品的id和购买的数量

	//开启事务，批量扣减商品的真实库存和锁定库存
	//提交事务
	return nil
}
