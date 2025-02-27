package model

import (
	"github.com/bytedance/sonic"
	"strconv"
)

type DelayMessage struct {
	Level int8   `json:"level"`
	Topic string `json:"topic"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (d *DelayMessage) String() string {
	return "DelayMessage{" +
		"level=" + strconv.Itoa(int(d.Level)) +
		", topic='" + d.Topic + "'" +
		", key='" + d.Key + "'" +
		", value='" + d.Value + "'" +
		"}"
}

func (d *DelayMessage) ToJsonBytes() []byte {
	bytes, err := sonic.Marshal(d)
	if err != nil {
		return nil
	}
	return bytes
}
