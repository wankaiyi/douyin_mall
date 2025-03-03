package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	Name string `gorm:"not null;type:varchar(100)" json:"name"`
	Base
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

func CategoryKey() string {
	return "categories"
}

func CategoryLockKey() string {
	return "categories:lock"
}
func GetAllCategoryByDb(db *gorm.DB, ctx context.Context, categories *[]*Category) (err error) {
	return db.WithContext(ctx).Model(&Category{}).Find(&categories).Error
}
