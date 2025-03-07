package producer

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/payment/conf"
	"douyin_mall/payment/infra/kafka/constant"
	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"

	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

var (
	paymentSuccessProducer sarama.SyncProducer
	err1                   error
)

func InitNormalPaymentProducer() {
	config := sarama.NewConfig()
	// 保证消息不丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Idempotent = true
	config.Producer.Retry.Max = 3
	config.Producer.Transaction.ID = uuid.New().String()
	config.Net.MaxOpenRequests = 1

	paymentSuccessProducer, err1 = sarama.NewSyncProducer(conf.GetConf().Kafka.BizKafka.BootstrapServers, config)
	if err1 != nil {
		panic(err1.Error())
	}

	server.RegisterShutdownHook(func() {
		_ = paymentSuccessProducer.Close()
	})

}

func SendPaymentSuccessOrderIdMsg(ctx context.Context, orderId string) {
	err = paymentSuccessProducer.BeginTxn()
	if err != nil {
		klog.Errorf("开启事务失败: %v", err)
		return
	}
	data, _ := sonic.Marshal(orderId)
	err = sendMessage(ctx, constant.PaymentSuccessTopic, data, constant.PaymentSuccessKeyPrefix+orderId)
	if err != nil {
		paymentSuccessProducer.AbortTxn()
		return
	}
	err = paymentSuccessProducer.CommitTxn()
	if err != nil {
		klog.Errorf("提交事务失败: %v", err)
		return
	}
}

func sendMessage(ctx context.Context, topic string, message []byte, key string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	otel.GetTextMapPropagator().Inject(ctx, tracing.NewProducerMessageCarrier(msg))

	partition, offset, err := paymentSuccessProducer.SendMessage(msg)
	if err != nil {
		klog.Errorf("消息发送失败, topic: %s, key: %s, value: %s, error: %v", topic, key, message, err)
		return err
	}
	klog.Infof("消息发送成功, topic: %s, key: %s, value: %s, partition: %d, offset: %d", topic, key, message, partition, offset)
	return nil
}
