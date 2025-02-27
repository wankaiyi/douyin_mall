package model

type UpdateProductSendToKafka struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Picture     string  `json:"picture,omitempty"`
}
