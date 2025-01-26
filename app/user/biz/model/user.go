package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `gorm:"unique"`
	Password string
}

func (u User) TableName() string {
	return "tb_user"
}

func GetUserByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	result := db.Create(user)
	return result.Error
}

func GetUserById(db *gorm.DB, ctx context.Context, userId int32) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Base: Base{ID: userId}}).First(&user).Error
	return
}
