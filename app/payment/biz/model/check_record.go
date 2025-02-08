package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// `gorm:"not null;type:varchar(64);uniqueIndex:idx_username_deleted_at"`
type CheckRecord struct {
	Base
	ReconDate     time.Time `gorm:"not null;type:datetime;"`
	AlipayTradeNo string    `gorm:"not null;type:varchar(64);uninqueIndex:idx_alipay_trade_no_deleted_at"`
	OrderId       string    `gorm:"not null;type:varchar(64);uninqueIndex:idx_order_id_deleted_at"`
	AlipayAmount  float64   `gorm:"not null;type:decimal(10,2);"`
	LocalAmount   float64   `gorm:"not null;type:decimal(10,2);"`
	Status        int       `gorm:"not null;type:int;comment:0一致，1不一致"`
}

func (CheckRecord) TableName() string {
	return "tb_check_records"
}
func CreateCheckRecord(db *gorm.DB, ctx context.Context, checkRecord *CheckRecord) (err error) {
	err = db.WithContext(ctx).Model(&CheckRecord{}).Create(checkRecord).Error
	return err
}
