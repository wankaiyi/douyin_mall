package kafka

import "douyin_mall/user/infra/kafka/consumer"

func Init() {
	consumer.InitUserCacheMessageConsumer()
}
