package producer

import (
	"context"
	"douyin_mall/payment/conf"
	"douyin_mall/payment/infra/kafka/constant"
	"github.com/bytedance/sonic"

	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

var (
	paymentSuccessProducer sarama.AsyncProducer
	err1                   error
)

func InitNormalPaymentProducer() {
	config := sarama.NewConfig()
	// 保证消息不丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	paymentSuccessProducer, err1 = sarama.NewAsyncProducer(conf.GetConf().Kafka.BizKafka.BootstrapServers, config)
	if err1 != nil {
		panic(err1.Error())
	}

	go func() {
		for msg := range paymentSuccessProducer.Successes() {
			klog.Infof("消息发送成功 消息内容: %s topic:%s partition:%d offset:%d\n", msg.Value, msg.Topic, msg.Partition, msg.Offset)
		}
	}()

	go func() {
		for err1 = range paymentSuccessProducer.Errors() {
			klog.Errorf("消息发送失败: %v\n", err1)
		}
	}()

	server.RegisterShutdownHook(func() {
		_ = paymentSuccessProducer.Close()
	})

}

func SendPaymentSuccessOrderIdMsg(ctx context.Context, orderId string) {
	data, _ := sonic.Marshal(orderId)
	sendMessage(ctx, constant.PaymentSuccessTopic, data, constant.PaymentSuccessKeyPrefix+orderId)
}
