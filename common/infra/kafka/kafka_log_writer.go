package kafka

import (
	"douyin_mall/common/utils/env"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"os/signal"
	"syscall"
)

type KafkaWriter struct {
	producer sarama.AsyncProducer
	topicId  string
}

func NewKafkaWriter(user, password, topicId string) *KafkaWriter {
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

	producer, err := sarama.NewAsyncProducer([]string{"gz-producer.cls.tencentcs.com:9096"}, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for err := range producer.Errors() {
			klog.Errorf("发送消息失败: %v", err)
		}
	}()

	// 捕获退出信号，优雅停机
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		_ = producer.Close()
	}()
	return &KafkaWriter{
		producer: producer,
		topicId:  topicId,
	}
}

// Write 实现 zapcore.WriteSyncer 的 Write 方法
func (kw *KafkaWriter) Write(p []byte) (n int, err error) {
	if currentEnv, err := env.GetString("env"); err != nil && currentEnv != "dev" {
		kw.producer.Input() <- &sarama.ProducerMessage{
			Topic: kw.topicId,
			Value: sarama.StringEncoder(p),
		}
	}
	return len(p), nil
}

// Sync 实现 zapcore.WriteSyncer 的 Sync 方法
func (kw *KafkaWriter) Sync() error {
	return nil
}
