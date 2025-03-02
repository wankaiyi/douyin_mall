package kafka

import (
	"douyin_mall/order/infra/kafka/consumer"
	"douyin_mall/order/infra/kafka/producer"
)

func Init() {
	consumer.InitDelayOrderConsumer()
	producer.InitDelayOrderProducer()
	producer.InitCancelOrderSuccessProducer()
}
