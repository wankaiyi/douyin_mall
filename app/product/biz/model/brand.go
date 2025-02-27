package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Brand struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Brand) TableName() string {
	return "tb_brand"
}

func SelectBrand(db *gorm.DB, ctx context.Context, id int64) (brand Brand, err error) {
	result := db.WithContext(ctx).Where("id=?", id).First(&brand)
	err = result.Error
	return brand, err
}

func CreateBrand(db *gorm.DB, ctx context.Context, brand *Brand) (err error) {
	result := db.WithContext(ctx).Create(&brand)
	err = result.Error
	return err
}

func DeleteBrand(db *gorm.DB, ctx context.Context, id int64) (err error) {
	result := db.WithContext(ctx).Where("id=?", id).Delete(&Brand{})
	err = result.Error
	return err
}

func UpdateBrand(db *gorm.DB, ctx context.Context, brand *Brand) (err error) {
	result := db.WithContext(ctx).Updates(&brand)
	err = result.Error
	return err
}
