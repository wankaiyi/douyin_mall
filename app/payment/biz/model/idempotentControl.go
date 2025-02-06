package model

import (
	"context"
	"gorm.io/gorm"
)

// IdempotentControl 支付宝异步回调的幂等控制表
type IdempotentControl struct {
	Base
	OrderID string `gorm:"not null;type:varchar(64);uniqueIndex:idx_order_id;comment:唯一索引的幂等健(订单id)" json:"order_id"`
}

func (IdempotentControl) TableName() string {
	return "tb_idempotent_control"
}
func AddIdempotentControl(db *gorm.DB, ctx context.Context, orderID string) (err error) {
	err = db.Create(&IdempotentControl{OrderID: orderID}).Error
	return
}
