package key

import (
	"hotkey/model/base"
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
		", ServiceName: " + h.ServiceName +
		", Count: " + h.Count.String() +
		", KeyRule: " + h.KeyRule.String() +
		", Remove: " + strconv.FormatBool(h.Remove) +
		"}"
}
