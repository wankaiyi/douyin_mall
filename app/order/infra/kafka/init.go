package kafka

import (
	"douyin_mall/order/infra/kafka/consumer"
	"douyin_mall/order/infra/kafka/producer"
)

func Init() {
	consumer.InitDelayOrderConsumer()
	consumer.InitDelayStatusCompensationConsumer()
	producer.InitDelayOrderProducer()
	producer.InitCancelOrderSuccessProducer()
}
