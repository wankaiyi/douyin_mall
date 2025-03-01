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
	cancelOrderProducer sarama.AsyncProducer
	err2                error
)

func InitCancelOrderProducer() {
	config := sarama.NewConfig()
	// 保证消息不丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	cancelOrderProducer, err = sarama.NewAsyncProducer(conf.GetConf().Kafka.BizKafka.BootstrapServers, config)
	if err2 != nil {
		panic(err.Error())
	}

	go func() {
		for msg := range cancelOrderProducer.Successes() {
			klog.Infof("消息发送成功 消息内容: %s topic:%s partition:%d offset:%d\n", msg.Value, msg.Topic, msg.Partition, msg.Offset)
		}
	}()

	go func() {
		for err2 = range cancelOrderProducer.Errors() {
			klog.Errorf("消息发送失败: %v\n", err2)
		}
	}()

	server.RegisterShutdownHook(func() {
		_ = cancelOrderProducer.Close()
	})

}

func SendCancelOrderMsg(ctx context.Context, orderId string) {
	data, _ := sonic.Marshal(orderId)
	sendMessage(ctx, constant.CancelOrderTopic, data, constant.CancelOrderKeyPrefix+orderId)
}
