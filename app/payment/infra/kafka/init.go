package kafka

import (
	"douyin_mall/payment/infra/kafka/consumer"
	"douyin_mall/payment/infra/kafka/producer"
)

func Init() {
	consumer.InitDelayOrderConsumer()
	producer.InitDelayCheckoutProducer()
	producer.InitNormalPaymentProducer()
	producer.InitCancelOrderProducer()
}
