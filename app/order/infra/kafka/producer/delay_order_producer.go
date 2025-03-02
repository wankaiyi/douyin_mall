package producer

import (
	"context"
	"douyin_mall/common/infra/kafka/model"
	"douyin_mall/order/conf"
	"douyin_mall/order/infra/kafka/constant"
	model2 "douyin_mall/order/infra/kafka/model"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

var (
	producer sarama.AsyncProducer
)

func InitDelayOrderProducer() {
	config := sarama.NewConfig()
	// 保证消息不丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err = sarama.NewAsyncProducer(conf.GetConf().Kafka.BizKafka.BootstrapServers, config)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for msg := range producer.Successes() {
			klog.Infof("延时取消订单消息发送成功 消息内容: %s topic:%s partition:%d offset:%d\n", msg.Value, msg.Topic, msg.Partition, msg.Offset)
		}
	}()

	go func() {
		for err = range producer.Errors() {
			klog.Errorf("延时取消订单消息发送失败: %v\n", err)
		}
	}()

	server.RegisterShutdownHook(func() {
		_ = producer.Close()
	})

}

func SendDelayOrder(ctx context.Context, orderId string, delayLevel int8) {
	delayOrderMessage := model2.DelayOrderMessage{OrderID: orderId, Level: delayLevel}

	delayMsg := &model.DelayMessage{
		Level: delayLevel,
		Topic: constant.DelayCancelOrderTopic,
		Key:   orderId,
		Value: delayOrderMessage.ToJson(),
	}
	delayMsgBytes, _ := sonic.Marshal(delayMsg)
	sendMessageAsync(ctx, constant.DelayTopic, delayMsgBytes, orderId, producer)
}
