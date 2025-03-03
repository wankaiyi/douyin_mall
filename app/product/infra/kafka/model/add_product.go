package model

type AddProductSendToKafka struct {
	ID          int32   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Picture     string  `json:"picture,omitempty"`
	Stock       int32   `json:"stock,omitempty"`
	LockStock   int32   `json:"lock_stock,omitempty"`
}
