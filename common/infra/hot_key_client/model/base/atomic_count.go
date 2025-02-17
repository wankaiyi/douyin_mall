package base

import (
	"strconv"
	"sync/atomic"
)

type AtomicCount struct {
	Count int64
}

func (ac *AtomicCount) Add(count int64) {
	atomic.AddInt64(&ac.Count, count)

}
func (ac *AtomicCount) GetCount() int64 {
	return atomic.LoadInt64(&ac.Count)
}
func (ac *AtomicCount) String() string {
	return "AtomicCount{" + "Count=" + strconv.FormatInt(ac.GetCount(), 10) + "}"
}

func (ac *AtomicCount) Increment() {
	atomic.AddInt64(&ac.Count, 1)
}
