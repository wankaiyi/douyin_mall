package vo

type ProductRedisDataVo struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	ID          int64   `json:"id,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Picture     string  `json:"picture,omitempty"`
}
