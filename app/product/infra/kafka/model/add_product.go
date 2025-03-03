package model

type AddProductSendToKafka struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Picture     string  `json:"picture,omitempty"`
	Stock       int64   `json:"stock,omitempty"`
	LockStock   int64   `json:"lock_stock,omitempty"`
}
