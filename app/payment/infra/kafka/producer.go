package kafka

import (
	"context"
	"douyin_mall/payment/conf"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

func GetProducer() *sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待所有副本确认
	config.Producer.Retry.Max = 5                             // 重试次数
	config.Producer.Return.Successes = true                   // 返回成功信息
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区

	// 连接 Kafka Broker
	bizKafka := conf.GetConf().Kafka.BizKafka
	brokers := strings.Split(bizKafka.BootstrapServers, ",")
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	return &producer
}

// SendDelayMsg 发送延迟消息
// 时间为3秒
func SendDelayMsg(msg *sarama.ProducerMessage, producer sarama.SyncProducer) {

	// 构造消息

	// 发送消息
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	klog.CtxInfof(context.Background(), "生产者发送了一条信息，Sent message to partition %d at offset %d\n", partition, offset)
}
