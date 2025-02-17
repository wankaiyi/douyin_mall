package util

import (
	"hotkey/model/key"
)

var (
	BlockingQueueSize = 1000000
	BlQueue           *BlockingQueue
)

func init() {
	BlQueue = NewBlockingQueue(BlockingQueueSize)
}

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
