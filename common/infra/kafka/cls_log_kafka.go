package kafka

import (
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"os/signal"
	"syscall"
)

var (
	producer sarama.AsyncProducer
	err      error
	// 日志主题id
	logTopicId string
)

func InitClsLogKafka(user string, password string, topicId string) {
	logTopicId = topicId
	config := sarama.NewConfig()

	config.Net.SASL.Mechanism = "PLAIN"
	config.Net.SASL.Version = int16(1)
	config.Net.SASL.Enable = true
	config.Net.SASL.User = user
	config.Net.SASL.Password = password
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = 0
	config.Version = sarama.V1_1_0_0
	config.Producer.Compression = sarama.CompressionGZIP

	producer, err = sarama.NewAsyncProducer([]string{"gz-producer.cls.tencentcs.com:9096"}, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for err := range producer.Errors() {
			klog.Error("发送消息失败: %v\n", err)
		}
	}()

	go func() {
		for success := range producer.Successes() {
			klog.Info("发送消息成功，topic %s, partition %d, offset %d\n", success.Topic, success.Partition, success.Offset)
		}
	}()

	// 捕获退出信号，优雅停机
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		_ = producer.Close()
	}()
}

func SendLogMessage(log string) {
	producer.Input() <- &sarama.ProducerMessage{
		Topic: logTopicId,
		Value: sarama.StringEncoder(log),
	}
	_ = producer.Close()
}
