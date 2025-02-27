package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Category) TableName() string {
	return "tb_category"
}

func SelectCategory(db *gorm.DB, ctx context.Context, id int64) (category Category, err error) {
	result := db.WithContext(ctx).Where("id=?", id).First(&category)
	err = result.Error
	return category, err
}

func CreateCategory(db *gorm.DB, ctx context.Context, category *Category) (err error) {
	result := db.WithContext(ctx).Create(&category)
	err = result.Error
	return err
}

func DeleteCategory(db *gorm.DB, ctx context.Context, id int64) (err error) {
	result := db.WithContext(ctx).Where("id=?", id).Delete(&Category{})
	err = result.Error
	return err
}

func UpdateCategory(db *gorm.DB, ctx context.Context, category *Category) (err error) {
	result := db.WithContext(ctx).Updates(&category)
	err = result.Error
	return err
}
