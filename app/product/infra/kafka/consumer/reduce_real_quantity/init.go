package reduce_real_quantity

import (
	"context"
	"douyin_mall/product/conf"
	"douyin_mall/product/infra/kafka/constant"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

func ReleaseRealQuantityConsumer() {
	config := sarama.NewConfig()
	strategies := make([]sarama.BalanceStrategy, 1)
	strategies[0] = sarama.NewBalanceStrategyRange()
	config.Consumer.Group.Rebalance.GroupStrategies = strategies
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	// 创建消费者
	brokers := strings.Split(conf.GetConf().Kafka.BizKafka.BootstrapServers, ",")
	groupId := "reduce-product-stock"
	if conf.GetConf().Env == "dev" {
		groupId += "_dev"
	}
	consumer, err := sarama.NewConsumerGroup(brokers, groupId, config)
	handler := ReduceRealQuantityHandler{}
	for {
		err = consumer.Consume(
			context.Background(),
			[]string{constant.PaymentSuccess},
			handler,
		)
		if err != nil {
			klog.Error("Error from consumer: ", err)
		}
	}
}
