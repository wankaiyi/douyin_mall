package utils

import (
	"douyin_mall/common/infra/hot_key_client/model/key"
)

type BlockingQueue struct {
	queue chan key.HotkeyModel
}

func NewBlockingQueue(size int) *BlockingQueue {
	return &BlockingQueue{
		queue: make(chan key.HotkeyModel, size),
	}
}
func (q *BlockingQueue) Put(item key.HotkeyModel) {
	q.queue <- item
}
func (q *BlockingQueue) Take() key.HotkeyModel {
	return <-q.queue
}
