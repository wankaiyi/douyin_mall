package consumer

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	"douyin_mall/order/conf"
	"douyin_mall/order/infra/kafka/constant"
	model2 "douyin_mall/order/infra/kafka/model"
	"douyin_mall/order/infra/kafka/producer"
	"douyin_mall/order/infra/rpc"
	"douyin_mall/rpc/kitex_gen/payment"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type msgConsumerGroup struct{}

func (msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := sess.Context()
	for msg := range claim.Messages() {
		topic := msg.Topic
		klog.Infof("收到消息，topic:%q partition:%d offset:%d  value:%s\n", topic, msg.Partition, msg.Offset, string(msg.Value))

		var delayCancelOrderMessage model2.DelayOrderMessage
		err := sonic.Unmarshal(msg.Value, &delayCancelOrderMessage)
		if err != nil {
			klog.Errorf("解析消息失败，topic:%q partition:%d offset:%d  value:%s\n", topic, msg.Partition, msg.Offset, string(msg.Value))
			sess.MarkMessage(msg, "")
			sess.Commit()
			continue
		}

		msgCtx := otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))
		_, span := otel.Tracer("delay-order-consumer").Start(msgCtx, "consume-delay-order-message")
		otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))

		orderId := delayCancelOrderMessage.OrderID
		unpaid, err := checkOrderUnpaid(ctx, orderId)
		if err != nil {
			span.End()
			continue
		}
		if unpaid {
			// 继续发消息或取消订单
			switch delayCancelOrderMessage.Level {
			case constant.DelayTopic1mLevel:
				{
					producer.SendDelayOrder(ctx, orderId, constant.DelayTopic4mLevel)
				}
			case constant.DelayTopic4mLevel:
				{
					producer.SendDelayOrder(ctx, orderId, constant.DelayTopic5mLevel)
				}
			case constant.DelayTopic5mLevel:
				{
					// 订单超时，取消订单，有兜底方案，不管结果；取消订单是幂等的，不考虑重复消费
					cancelOrder(ctx, orderId)
				}
			}
		}

		span.End()
		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}

func cancelOrder(ctx context.Context, orderId string) {
	// 先取消支付，再取消订单
	err := cancelCharge(ctx, orderId)
	if err != nil {
		klog.Errorf("取消支付失败，orderId:%s，err:%v", orderId, err)
		return
	}

	affectedRows, err := model.CancelOrder(ctx, mysql.DB, orderId)
	if err != nil {
		klog.Errorf("取消订单失败，orderId:%s，err:%v", orderId, err)
		return
	}
	if affectedRows == 0 {
		// 主要是防止有支付成功回调请求并发，导致并发安全问题
		klog.Infof("订单状态不是未支付，无法取消订单，orderId:%s", orderId)
	}
	klog.Info("订单取消成功，orderId:%s", orderId)
	producer.SendCancelOrderSuccessMessage(ctx, orderId)
}

func cancelCharge(ctx context.Context, orderId string) error {
	req := &payment.CancelChargeReq{
		OrderId: orderId,
	}
	cancelChargeResp, err := rpc.PaymentClient.CancelCharge(ctx, req)
	if err != nil {
		return err
	}
	klog.Info("rpc调用取消支付，req:%v, resp:%v", req, cancelChargeResp)
	if cancelChargeResp.StatusCode != 0 {
		return errors.New("取消支付失败")
	}
	return nil
}

func checkOrderUnpaid(ctx context.Context, orderId string) (bool, error) {
	unpaid, err := model.CheckOrderUnpaid(ctx, mysql.DB, orderId)
	if err != nil {
		klog.Errorf("查询订单状态失败，orderId:%s，err:%v", orderId, err)
		return false, err
	}
	return unpaid, nil
}

func InitDelayOrderConsumer() {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V3_5_0_0
	consumerConfig.Consumer.Offsets.AutoCommit.Enable = false
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerConfig.Consumer.Offsets.Retry.Max = 3

	groupId := constant.DelayCancelOrderGroupId
	if conf.GetConf().Env == "dev" {
		groupId += "-dev"
	}
	cGroup, err := sarama.NewConsumerGroup(conf.GetConf().Kafka.BizKafka.BootstrapServers, groupId, consumerConfig)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err = cGroup.Consume(
				context.Background(),
				[]string{constant.DelayCancelOrderTopic},
				msgConsumerGroup{},
			)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}
	}()

	server.RegisterShutdownHook(func() {
		_ = cGroup.Close()
	})

}
