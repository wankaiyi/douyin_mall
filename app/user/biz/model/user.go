package model

import (
	"context"
	"douyin_mall/user/biz/dal/redis"
	redisUtils "douyin_mall/user/utils/redis"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	Base
	Username    string                `gorm:"not null;type:varchar(64);uniqueIndex:idx_username_deleted_at"`
	Email       string                `gorm:"not null;type:varchar(64)"`
	Sex         int32                 `gorm:"not null;type:tinyint;comment:性别：0-未知，1-男，2-女;default:0"`
	Password    string                `gorm:"not null;type:varchar(255)"`
	Description string                `gorm:"not null;type:varchar(255);default:''"`
	Avatar      string                `gorm:"not null;type:varchar(255);default:''"`
	Role        Role                  `gorm:"not null;type:varchar(64);default:'user'"`
	DeletedAt   soft_delete.DeletedAt `gorm:"index;uniqueIndex:idx_username_deleted_at"`
}

type UserBasicInfo struct {
	Username    string
	Email       string
	Sex         int32
	Description string
	Avatar      string
	CreatedAt   string
}

func (u UserBasicInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"username":    u.Username,
		"email":       u.Email,
		"sex":         SexToString(u.Sex),
		"description": u.Description,
		"avatar":      u.Avatar,
		"created_at":  u.CreatedAt,
	}
}

func (u User) TableName() string {
	return "tb_user"
}

func GetUserByUsername(db *gorm.DB, ctx context.Context, username string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Username: username}).First(&user).Error
	return
}

func CreateUser(db *gorm.DB, ctx context.Context, user *User) error {
	result := db.WithContext(ctx).Create(user)
	return result.Error
}

func GetUserById(db *gorm.DB, ctx context.Context, userId int32) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Base: Base{ID: userId}}).First(&user).Error
	return
}

func GetUserBasicInfoById(db *gorm.DB, ctx context.Context, userId int32) (user *UserBasicInfo, err error) {
	err = db.WithContext(ctx).Model(&User{}).Select("username", "email", "sex", "description", "avatar", "created_at").Where(&User{Base: Base{ID: userId}}).First(&user).Error
	return
}

func SexToString(sex int32) string {
	switch sex {
	case 0:
		return "未知"
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "未知"
	}
}

func UpdateUser(db *gorm.DB, ctx context.Context, id int32, username string, email string, sex int32, description string, avatar string) error {
	user := User{Base: Base{ID: id}}
	return db.WithContext(ctx).Model(&user). /*Save(&user).*/ Select("username", "email", "sex", "description", "avatar").Updates(
		User{
			Username:    username,
			Email:       email,
			Sex:         sex,
			Description: description,
			Avatar:      avatar,
		},
	).Error
}

func DeleteUser(db *gorm.DB, id int32) error {
	return db.Delete(&User{Base: Base{ID: id}}).Error
}

func GetUserRoleById(ctx context.Context, db *gorm.DB, id int32) (role string, err error) {
	var user *User
	err = db.WithContext(ctx).Select("role").Where(User{Base: Base{ID: id}}).First(&user).Error
	return string(user.Role), err
}

func CacheUserInfo(ctx context.Context, userInfo *UserBasicInfo, userId int32) error {
	key := redisUtils.GetUserKey(userId)
	err := redis.RedisClient.HSet(ctx, key, userInfo.ToMap()).Err()
	if err != nil {
		return err
	}
	// 设置过期时间和access token的过期时间一致
	err = redis.RedisClient.Expire(ctx, key, time.Hour*2).Err()
	if err != nil {
		return err
	}
	return nil
}
