package kafka

import (
	"context"
	"douyin_mall/product/conf"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
	"sync"
)

var (
	Producer sarama.AsyncProducer
	Consumer sarama.Consumer
	Topic    string
	once     sync.Once
	err      error
)

func InitClient() {
	// 配置Topic
	Topic = conf.GetConf().Kafka.BizKafka.ProductTopicId
	once.Do(func() {
		// 配置生产者
		err = InitProducer()
		if err != nil {
			return
		}
		// 配置消费者
		err = InitConsumer()
		if err != nil {
			return
		}

	})
}

func InitProducer() (err error) {
	config := sarama.NewConfig()
	// 配置生产者
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = 0
	config.Version = sarama.V1_1_0_0
	config.Producer.Compression = sarama.CompressionGZIP

	// 创建生产者
	brokers := strings.Split(conf.GetConf().Kafka.BizKafka.BootstrapServers, ",")
	Producer, err = sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		klog.Error("Failed to start producer:", err)
	}
	klog.Info("Successfully connected to kafka", Producer)
	return
}

func InitConsumer() (err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V1_1_0_0
	// 创建消费者
	brokers := strings.Split(conf.GetConf().Kafka.BizKafka.BootstrapServers, ",")
	group, err := sarama.NewConsumerGroup(brokers, "product_group", config)
	err = group.Consume(
		context.Background(),
		[]string{conf.GetConf().Kafka.BizKafka.ProductTopicId},
		ProductKafkaConsumer{},
	)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	Consumer, err = sarama.NewConsumer(brokers, config)
	if err != nil {
		klog.Error("Failed to start consumer:", err)
	}
	klog.Info("Successfully connected to kafka", Consumer)
	return
}
