package producer

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel"
)

var (
	err error
)

func sendMessage(ctx context.Context, topic string, message []byte, key string, producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	otel.GetTextMapPropagator().Inject(ctx, tracing.NewProducerMessageCarrier(msg))

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	klog.Infof("消息发送成功, topic: %s, partition: %d, offset: %d", topic, partition, offset)
	return nil
}

func sendMessageAsync(ctx context.Context, topic string, message []byte, key string, producer sarama.AsyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	otel.GetTextMapPropagator().Inject(ctx, tracing.NewProducerMessageCarrier(msg))

	producer.Input() <- msg
}
