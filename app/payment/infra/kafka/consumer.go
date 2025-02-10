package kafka

import (
	"context"
	"douyin_mall/payment/biz/dal/alipay"
	"douyin_mall/payment/conf"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"strings"
)

type ConsumerGroupHandler struct{}

func (h ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		klog.CtxInfof(context.Background(), "消费者接受到消息，Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s \n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		session.MarkMessage(msg, "done")
		//todo 查询订单状态是否为已支付，如果是，则更新订单状态

		//不是已支付状态，查询支付宝
		orderId, err := strconv.ParseInt(string(msg.Value), 0, 64)
		if err != nil {
			klog.CtxErrorf(context.Background(), "解析订单号失败，err=%v", err)
			return err
		}
		tradeStatus, err := alipay.QueryOrder(context.Background(), orderId)
		if err != nil {
			klog.CtxErrorf(context.Background(), "查询支付宝订单状态失败，err=%v", err)
			return err
		}
		if tradeStatus != "TRADE_SUCCESS" {
			klog.CtxInfof(context.Background(), "支付宝订单状态不为TRADE_SUCCESS,orderId=%d,tradeStatus=%s , ", orderId, tradeStatus)

			return nil
		}
		//todo 更新订单状态

	}
	return nil
}
func ConsumerGroupInit() {
	config := sarama.NewConfig()
	strategies := make([]sarama.BalanceStrategy, 1)
	strategies[0] = sarama.NewBalanceStrategyRange()
	config.Consumer.Group.Rebalance.GroupStrategies = strategies
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	bizKafka := conf.GetConf().Kafka.BizKafka
	brokers := strings.Split(bizKafka.BootstrapServers, ",")
	consumer, err := sarama.NewConsumerGroup(brokers, "check-group-test", config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	handler := ConsumerGroupHandler{}
	for {
		err := consumer.Consume(context.Background(), []string{"check-topic-test"}, &handler)
		if err != nil {
			panic(err)
		}
	}
}
