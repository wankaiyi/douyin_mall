package kafka

import (
	producer "douyin_mall/auth/infra/kafka/producer"
)

func Init() {
	producer.InitUserCacheMessageProducer()
}
