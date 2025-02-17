package key

import "fmt"

type KeyRule struct {
	//触发间隔,单位为秒
	Interval int64 `json:"interval"`
	//触发阈值
	Threshold int64 `json:"threshold"`
	//在client端的缓存时间，单位为秒
	Duration int64 `json:"duration"`
}

func (k *KeyRule) String() string {
	return fmt.Sprintf("Interval: %d, Threshold: %d", k.Interval, k.Threshold)
}
