package model

import "github.com/bytedance/sonic"

type DelayStockCompensationMessage struct {
	Uuid         string        `json:"uuid"`
	ProductItems []ProductItem `json:"product_items"`
}

type ProductItem struct {
	ProductID int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

func JsonToObj(jsonStr string) *DelayStockCompensationMessage {
	var obj DelayStockCompensationMessage
	_ = sonic.Unmarshal([]byte(jsonStr), &obj)
	return &obj
}

func (d *DelayStockCompensationMessage) ToJson() string {
	jsonBytes, _ := sonic.Marshal(d)
	return string(jsonBytes)
}
