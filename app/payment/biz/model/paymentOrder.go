package model

import (
	"context"
	"gorm.io/gorm"
)

//`gorm:"not null;type:varchar(64);uniqueIndex:idx_username_deleted_at"`
//`gorm:"not null;type:varchar(64)"`
//`gorm:"not null;type:tinyint;common:性别：0-未知，1-男，2-女;default:0"`
//`gorm:"not null;type:varchar(255)"`
//`gorm:"not null;type:varchar(255);default:''"`
//`gorm:"not null;type:varchar(255);default:''"`
//`gorm:"not null;type:varchar(64);default:'user'"`
//`gorm:"index;uniqueIndex:idx_username_deleted_at"`

type PaymentOrder struct {
	Base
	OrderID string  `gorm:"not null;type:varchar(64);default:'';uniqueIndex:idx_order_id"`
	UserID  int32   `gorm:"not null;type: int;normalIndex:idx_user_id;"`
	Amount  float64 `gorm:"not null;type: decimal default 0.00;"`
	Status  int32   `gorm:"not null;type:tinyint;comment:订单状态：0-待支付，1-支付成功，2-支付失败"`
}

func (po *PaymentOrder) TableName() string {
	return "tb_payment_orders"
}
func GetPaymentOrdersByOrderID(db *gorm.DB, ctx context.Context, orderID string) (paymentOrders *PaymentOrder, err error) {
	result := db.WithContext(ctx).Model(&PaymentOrder{}).Where("order_id=?", orderID).First(&paymentOrders)
	if result.RowsAffected < 1 {
		return nil, result.Error
	}
	return
}
func AddPaymentOrders(db *gorm.DB, ctx context.Context, paymentOrders *PaymentOrder) error {
	return db.WithContext(ctx).Create(paymentOrders).Error
}
func UpdatePaymentOrders(db *gorm.DB, ctx context.Context, paymentOrders *PaymentOrder) error {
	return db.WithContext(ctx).Save(paymentOrders).Error
}
func DeletePaymentOrders(db *gorm.DB, ctx context.Context, paymentOrders *PaymentOrder) error {
	return db.WithContext(ctx).Delete(paymentOrders).Error
}
