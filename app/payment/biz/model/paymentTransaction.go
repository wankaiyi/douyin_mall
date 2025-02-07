package model

import "gorm.io/gorm"

type PaymentTransaction struct {
	Base
	OrderID       string `gorm:"not null;type:varchar(64);normalIndex:idx_order_id"`
	AlipayTradeNo string `gorm:"not null;type:varchar(64);uniqueIndex:idx_alipay_trade_no;comment:支付宝交易号"`
	TradeStatus   string `gorm:"not null;type:varchar(64);comment:交易状态"`
	Callback      string `gorm:"not null;type:text;comment:支付宝异步回调的原文"`
	RequestParams string `gorm:"not null;type:text;comment:调用支付宝接口的原始请求参数"`
}

func (PaymentTransaction) TableName() string {
	return "tb_payment_transactions"
}
func AddPaymentTransaction(db *gorm.DB, pt *PaymentTransaction) error {
	return db.Create(pt).Error
}
