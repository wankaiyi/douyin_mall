package model

import "github.com/bytedance/sonic"

type DelayOrderMessage struct {
	Level   int8   `json:"level"`
	OrderID string `json:"order_id"`
}

func JsonToObj(jsonStr string) *DelayOrderMessage {
	var obj DelayOrderMessage
	_ = sonic.Unmarshal([]byte(jsonStr), &obj)
	return &obj
}

func (d *DelayOrderMessage) ToJson() string {
	jsonBytes, _ := sonic.Marshal(d)
	return string(jsonBytes)
}
