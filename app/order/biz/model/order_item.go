package model

import (
	"context"
	"gorm.io/gorm"
)

type OrderItem struct {
	Base
	OrderID   string  `gorm:"not null;type:varchar(64);index:idx_order_id"`
	ProductID int32   `gorm:"not null;type:int"`
	Quantity  int32   `gorm:"not null;type:int"`
	Price     float64 `gorm:"not null;type:decimal(10,2)"`
	Cost      float64 `gorm:"not null;type:decimal(10,2)"`
}

func (o *OrderItem) TableName() string {
	return "tb_order_item"
}

func CreateOrderItems(ctx context.Context, db *gorm.DB, list []*OrderItem) error {
	return db.WithContext(ctx).Create(list).Error
}

func GetOrderItemsByOrderIdList(ctx context.Context, db *gorm.DB, list []string) ([]*OrderItem, error) {
	var orderItems []*OrderItem
	err := db.WithContext(ctx).Model(&OrderItem{}).Where("order_id IN ?", list).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}
