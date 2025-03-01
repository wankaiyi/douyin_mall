package release_real_quantity

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

	// 创建消费者
	brokers := strings.Split(conf.GetConf().Kafka.BizKafka.BootstrapServers, ",")
	groupId := "product_group_add"
	if conf.GetConf().Env == "dev" {
		groupId += "_dev"
	}
	consumer, err := sarama.NewConsumerGroup(brokers, groupId, config)
	handler := ReleaseRealQuantityHandler{}
	for {
		err = consumer.Consume(
			context.Background(),
			[]string{constant.AddTopic},
			handler,
		)
		if err != nil {
			klog.Error("Error from consumer: ", err)
		}
	}
}
