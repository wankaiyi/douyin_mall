package key

import (
	"douyin_mall/common/infra/hot_key_client/constants"
	"douyin_mall/common/infra/hot_key_client/model/base"
	"strconv"
)

type HotkeyModel struct {
	ID string `json:"id"`
	//创建时间时间戳，单位毫秒
	CreateAt    int64            `json:"create_at"`
	Key         string           `json:"key"`
	ServiceName string           `json:"service_name"`
	Count       base.AtomicCount `json:"atomic_count"`
	//是否删除
	Remove  bool `json:"remove"`
	KeyRule `json:"key_rule"`
}

func (h *HotkeyModel) String() string {
	return "HotkeyModel{" +
		"ID: " + h.ID +
		", CreateAt: " + strconv.Itoa(int(h.CreateAt)) +
		", Key: " + h.Key +
		", Count: " + h.Count.String() +
		", Remove: " + strconv.FormatBool(h.Remove) +
		", KeyRule: " + h.KeyRule.String() +
		"}"
}

// NewDefaultHotkeyModel1 创建一个默认的热键模型,规则是1秒内最多访问1次
func NewDefaultHotkeyModel1(key string) *HotkeyModel {
	var hotKeyModel HotkeyModel
	hotKeyModel.Key = key
	hotKeyModel.ServiceName = "test-service"
	hotKeyModel.Threshold = 1
	hotKeyModel.Duration = 3000
	hotKeyModel.Interval = 1
	return &hotKeyModel
}

// NewHotKeyModelWithConfig 创建一个热键模型，根据配置创建
func NewHotKeyModelWithConfig(key string, config *KeyConf) *HotkeyModel {
	var hotKeyModel HotkeyModel
	hotKeyModel.Key = key
	hotKeyModel.ServiceName = config.ServiceName
	hotKeyModel.Threshold = config.Threshold
	hotKeyModel.Duration = config.Duration
	hotKeyModel.Interval = config.Interval
	return &hotKeyModel

}
func NewDeleteHotkeyModel(key string) *HotkeyModel {
	var hotKeyModel HotkeyModel
	hotKeyModel.Key = key
	hotKeyModel.ServiceName = constants.ClientServiceName
	hotKeyModel.Remove = true
	return &hotKeyModel
}
