package vo

var (
	Add    string = "add"
	Update string = "update"
	Delete string = "delete"
)

type Type struct {
	Name string `json:"name,omitempty"`
}

type ProductSendToKafka struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
type ProductKafkaDataVO struct {
	Type    Type               `json:"type,omitempty"`
	Product ProductSendToKafka `json:"product,omitempty"`
}
