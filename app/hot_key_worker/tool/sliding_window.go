package tool

import (
	"hotkey/model/base"
	"sync"
	"time"
)

type SlidingWindow struct {
	//装多个窗口的队列，等于windowsSize*2
	TimeSlices *base.AtomicIntArray
	//队列的长度
	TimeSliceSize int64
	//每个时间槽的长度(时长)单位为毫秒
	TimeMillPerSlice int64
	//窗口大小
	WindowSize int64
	//阈值，超过阈值则触发事件
	Threshold int64
	//窗口创建时间戳
	BeginTimeStamp int64
	//窗口结束时间戳
	EndTimeStamp int64
	//读写锁
	Mutex *sync.Mutex
}

func NewSlidingWindow(duration int64, threshold int64) *SlidingWindow {
	if duration > 600 {
		duration = 600
	}
	var windowSize int64
	var timeMillPerSlice int64
	if duration < 5 {
		windowSize = 5
		timeMillPerSlice = duration * 200
	} else {
		windowSize = 10
		timeMillPerSlice = duration * 100
	}
	var slidingWindow = &SlidingWindow{
		TimeSliceSize:    windowSize * 2,
		TimeMillPerSlice: timeMillPerSlice,
		WindowSize:       windowSize,
		Threshold:        threshold,
		Mutex:            &sync.Mutex{},
	}
	reset(slidingWindow)
	return slidingWindow

}

// reset 窗口的初始化
func reset(slidingWindow *SlidingWindow) {
	slidingWindow.BeginTimeStamp = time.Now().UnixMilli()
	atomicIntArray := base.NewAtomicIntArray(slidingWindow.TimeSliceSize)
	var i int64
	for i = 0; i < slidingWindow.TimeSliceSize; i++ {
		atomicIntArray.Set(i, 0)
	}
	slidingWindow.TimeSlices = atomicIntArray

}
func (slidingWindow *SlidingWindow) LocationIndex() int64 {
	now := time.Now().UnixMilli()
	//当前时间不在窗口内，重置窗口
	if now-slidingWindow.EndTimeStamp > slidingWindow.TimeSliceSize*slidingWindow.TimeMillPerSlice {
		reset(slidingWindow)
	}
	//当前时间在窗口内，计算当前时间所在的槽位索引
	index := (now - slidingWindow.BeginTimeStamp) / slidingWindow.TimeMillPerSlice
	if index < 0 {
		return 0
	}
	return index
}

// AddCount 增加计数
func (slidingWindow *SlidingWindow) AddCount(count int64) bool {
	//加锁
	slidingWindow.Mutex.Lock()
	defer slidingWindow.Mutex.Unlock()

	index := slidingWindow.LocationIndex()
	slidingWindow.ClearFormIndex(index)
	var sum int64
	sum = 0
	var i int64
	sum += slidingWindow.TimeSlices.AddAndGet(index, count)
	for i = 1; i < slidingWindow.WindowSize; i++ {
		sum += slidingWindow.TimeSlices.Get((index - i + slidingWindow.TimeSliceSize) % slidingWindow.TimeSliceSize)
	}
	slidingWindow.EndTimeStamp = time.Now().UnixMilli()
	return sum >= slidingWindow.Threshold

}

// ClearFormIndex 当前时间片内windowSize到windowSize*2的窗口内的计数清零
func (slidingWindow *SlidingWindow) ClearFormIndex(index int64) {

	for i := 1; i < int(slidingWindow.WindowSize); i++ {
		j := index + int64(i)
		if j >= slidingWindow.WindowSize*2 {
			j -= slidingWindow.WindowSize * 2
		}
		slidingWindow.TimeSlices.Set(j, 0)
	}

}
