package kafka

import (
	"douyin_mall/product/conf"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

var (
	Producer sarama.SyncProducer
	Topic    string
	err      error
)

func InitClient() {
	// 配置Topic
	Topic = conf.GetConf().Kafka.BizKafka.ProductTopicId
	// 配置生产者
	err = InitProducer()
	if err != nil {
		return
	}

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
	Producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		klog.Error("Failed to start producer:", err)
	}
	klog.Info("Successfully connected to kafka", Producer)
	return
}
