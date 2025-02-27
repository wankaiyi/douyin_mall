package add

import (
	"context"
	"douyin_mall/product/conf"
	"douyin_mall/product/infra/kafka/constant"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

func UpdateConsumer() {
	config := sarama.NewConfig()
	strategies := make([]sarama.BalanceStrategy, 1)
	strategies[0] = sarama.NewBalanceStrategyRange()
	config.Consumer.Group.Rebalance.GroupStrategies = strategies
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建消费者
	brokers := strings.Split(conf.GetConf().Kafka.BizKafka.BootstrapServers, ",")
	groupId := "product_group_update"
	if conf.GetConf().Env == "dev" {
		groupId += "_dev"
	}
	consumer, err := sarama.NewConsumerGroup(brokers, groupId, config)
	handler := UpdateProductHandler{}
	for {
		err = consumer.Consume(
			context.Background(),
			[]string{constant.UpdateTopic},
			handler,
		)
		if err != nil {
			klog.Error("Error from consumer: ", err)
		}
	}
}
