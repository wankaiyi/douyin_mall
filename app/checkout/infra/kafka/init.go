package kafka

import (
	"douyin_mall/checkout/infra/kafka/producer"
)

func Init() {
	producer.InitDelayStockCompensationProducer()
}
