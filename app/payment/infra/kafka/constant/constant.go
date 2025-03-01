package constant

const (
	DelayCheckOrderPrefix  = "delay_check_order_"
	DelayCheckOrderGroupId = "delay_check_order_group_"
	DelayCheckOrderTopic   = "delay-check-order"

	DelayTopic         = "delay-message"
	DelayTopic30sLevel = 3
	DelayTopic1mLevel  = 4
	DelayTopic3mLevel  = 6
	DelayTopic4mLevel  = 7
	DelayTopic5mLevel  = 8

	PaymentSuccessKeyPrefix = "payment_success_"
	PaymentSuccessTopic     = "payment-success"

	CancelOrderKeyPrefix = "cancel_order_"
	CancelOrderTopic     = "cancel-order"
)
