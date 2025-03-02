package consumer

import (
	add "douyin_mall/product/infra/kafka/consumer/add"
	del "douyin_mall/product/infra/kafka/consumer/delete"
	rel "douyin_mall/product/infra/kafka/consumer/release_real_quantity"
	upd "douyin_mall/product/infra/kafka/consumer/update"
)

func InitConsumer() {
	go add.AddConsumer()
	go del.DeleteConsumer()
	go upd.UpdateConsumer()
	go rel.ReleaseRealQuantityConsumer()
}
