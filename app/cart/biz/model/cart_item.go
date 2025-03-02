package model

import (
	"context"
	"gorm.io/gorm"
)

type CartItem struct {
	Base
	UserId    int32 `gorm:"not null;type:int;uniqueIndex:idx_user_id_product_id"`
	ProductId int32 `gorm:"not null;type:int;uniqueIndex:idx_user_id_product_id"`
	Quantity  int32 `gorm:"not null;type:int;"`
}

func (c CartItem) TableName() string {
	return "tb_cart_items"
}

func AddCartItem(ctx context.Context, db *gorm.DB, item *CartItem) error {
	var count int64
	err := db.WithContext(ctx).Model(&CartItem{}).
		Where(&CartItem{UserId: item.UserId, ProductId: item.ProductId}).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return db.WithContext(ctx).Model(&CartItem{}).
			Where(&CartItem{UserId: item.UserId, ProductId: item.ProductId}).
			Update("quantity", gorm.Expr("quantity + ?", item.Quantity)).Error
	} else {
		return db.WithContext(ctx).Create(item).Error
	}
}

func GetCartItemByUserId(ctx context.Context, db *gorm.DB, userId int32) ([]*CartItem, error) {
	var items []*CartItem
	err := db.WithContext(ctx).Where(&CartItem{UserId: userId}).Find(&items).Error
	return items, err
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId int32) error {
	err := db.WithContext(ctx).Where(&CartItem{UserId: userId}).Delete(&CartItem{}).Error
	return err
}
