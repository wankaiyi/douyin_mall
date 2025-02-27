package kafka

import (
	"douyin_mall/product/infra/kafka/consumer"
	"douyin_mall/product/infra/kafka/producer"
)

func Init() {
	consumer.InitConsumer()
	producer.InitProducerClient()
}
