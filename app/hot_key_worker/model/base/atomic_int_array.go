package base

import "sync/atomic"

type AtomicIntArray struct {
	Data []int64
}

func NewAtomicIntArray(size int64) *AtomicIntArray {
	return &AtomicIntArray{
		Data: make([]int64, size),
	}
}

func (arr *AtomicIntArray) AddAndGet(index int64, value int64) int64 {
	atomic.AddInt64(&arr.Data[index], value)
	return atomic.LoadInt64(&arr.Data[index])

}
func (arr *AtomicIntArray) Set(index int64, value int64) {
	atomic.StoreInt64(&arr.Data[index], value)
}
func (arr *AtomicIntArray) Get(index int64) int64 {
	return atomic.LoadInt64(&arr.Data[index])
}
