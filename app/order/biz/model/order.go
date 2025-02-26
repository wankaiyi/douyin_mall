package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const (
	OrderStatusUnpaid = iota
	OrderStatusPaid
	OrderStatusCanceled
)

type Order struct {
	OrderID       string  `gorm:"primarykey;type:varchar(64)"`
	UserID        int32   `gorm:"not null;type:int;index:idx_user_id"`
	TotalCost     float64 `gorm:"not null;type:decimal(10,2)"`
	Name          string  `gorm:"not null;type:varchar(64)"`
	PhoneNumber   string  `gorm:"not null;type:char(11)"`
	Province      string  `gorm:"not null;type:varchar(64)"`
	City          string  `gorm:"not null;type:varchar(64)"`
	Region        string  `gorm:"not null;type:varchar(64)"`
	DetailAddress string  `gorm:"not null;type:varchar(255)"`
	Status        int32   `gorm:"not null;type:int;default:0"`
	CreatedAt     time.Time
}

func (o *Order) TableName() string {
	return "tb_order"
}

func CreateOrder(ctx context.Context, db *gorm.DB, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}

func GetOrdersByUserId(ctx context.Context, db *gorm.DB, id int32) (orderList []Order, err error) {
	err = db.WithContext(ctx).Where(&Order{UserID: id}).Find(&orderList).Error
	return
}
