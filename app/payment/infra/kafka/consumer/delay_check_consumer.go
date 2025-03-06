package consumer

import (
	"context"
	commonConstant "douyin_mall/common/constant"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/payment/biz/dal/alipay"
	"douyin_mall/payment/conf"
	"douyin_mall/payment/infra/kafka/constant"
	model2 "douyin_mall/payment/infra/kafka/model"
	"douyin_mall/payment/infra/kafka/producer"
	"douyin_mall/payment/infra/rpc"
	"douyin_mall/rpc/kitex_gen/order"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"go.opentelemetry.io/otel"
	"strconv"
	"strings"
)

type msgConsumerGroup struct{}

func (msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := sess.Context()
	for msg := range claim.Messages() {
		topic := msg.Topic
		klog.Infof("收到消息，topic:%q partition:%d offset:%d  value:%s\n", topic, msg.Partition, msg.Offset, string(msg.Value))
		if !strings.HasPrefix(string(msg.Key), constant.DelayCheckOrderPrefix) {
			sess.MarkMessage(msg, "")
			sess.Commit()
			continue
		}

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
		id, _ := strconv.ParseInt(orderId, 10, 64)
		klog.CtxInfof(ctx, "查询支付宝系统订单状态,orderId=%s", orderId)
		status, err := checkPaymentStatus(ctx, id)

		if err != nil {
			span.End()
			continue
		}
		if !status {
			// 10分钟后订单会自动，所以要在10分钟时查一次，10分钟后查一次
			switch delayCancelOrderMessage.Level {
			case constant.DelayTopic1mLevel:
				{
					producer.SendCheckoutDelayMsg(ctx, orderId, constant.DelayTopic3mLevel)
				}
			case constant.DelayTopic4mLevel:
				{
					producer.SendCheckoutDelayMsg(ctx, orderId, constant.DelayTopic5mLevel)
				}
			case constant.DelayTopic5mLevel:
				{
					//过了10分钟，订单有可能被取消，也有可能成功支付了，所以要30s后再查一次
					producer.SendCheckoutDelayMsg(ctx, orderId, constant.DelayTopic30sLevel)
				}
			case constant.DelayTopic30sLevel:
				{
					klog.CtxInfof(ctx, "订单:orderId=%d, 超10分钟仍然未被支付，已被自动关闭！", orderId)
				}
			}
		}

		span.End()
		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}
func checkPaymentStatus(ctx context.Context, orderId int64) (bool, error) {
	// 查询订单状态是否为已支付，如果是，则更新订单状态
	orderResp, err2 := rpc.OrderClient.GetOrder(ctx, &order.GetOrderReq{
		OrderId: strconv.FormatInt(orderId, 10),
	})
	if err2 != nil {
		return false, err2
	}
	if orderResp.Order.Status == commonConstant.OrderStatusUnpaid {
		//不是已支付状态，查询支付宝
		tradeStatus, err := alipay.QueryOrder(context.Background(), orderId)
		if err != nil {
			return false, err
		}
		if tradeStatus == "TRADE_CLOSED" {
			klog.CtxInfof(context.Background(), "支付宝订单状态为交易已关闭,orderId=%d,tradeStatus=%s , ", orderId, tradeStatus)
			return true, nil
		}
		if tradeStatus != "TRADE_SUCCESS" {
			klog.CtxInfof(context.Background(), "支付宝订单状态不为TRADE_SUCCESS,orderId=%d,tradeStatus=%s , ", orderId, tradeStatus)
			return false, nil
		}
		//支付成功，更新订单状态
		klog.CtxInfof(context.Background(), "支付宝订单状态为TRADE_SUCCESS,orderId=%d,tradeStatus=%s , 更新订单状态", orderId, tradeStatus)
		markOrderPaidResp, err := rpc.OrderClient.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{
			OrderId: strconv.FormatInt(orderId, 10),
		})
		if err != nil {
			return false, err
		}
		if markOrderPaidResp.StatusCode != 0 {
			klog.CtxErrorf(context.Background(), "更新订单状态失败,orderId=%d,status_code=%d,status_msg=%s", orderId, markOrderPaidResp.StatusCode, markOrderPaidResp.StatusMsg)
			return false, nil
		}

		return true, nil
	}
	return false, nil

}

func InitDelayOrderConsumer() {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V3_5_0_0
	consumerConfig.Consumer.Offsets.AutoCommit.Enable = false
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerConfig.Consumer.Offsets.Retry.Max = 3

	groupId := constant.DelayCheckOrderGroupId
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
				[]string{constant.DelayCheckOrderTopic},
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
