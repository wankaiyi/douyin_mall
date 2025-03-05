package model

import (
	"context"
	"gorm.io/gorm"
)

const (
	AddressDefaultStatusNo  = 0
	AddressDefaultStatusYes = 1
)

type Address struct {
	Base
	UserId        int32  `gorm:"not null;type:int;index:idx_user_id"`
	Name          string `gorm:"not null;type:varchar(64);comment:收件人姓名"`
	PhoneNumber   string `gorm:"not null;type:varchar(64);comment:收件人手机号"`
	DefaultStatus int8   `gorm:"not null;type:int(1);default:0;comment:是否默认地址，0-否，1-是"`
	Province      string `gorm:"not null;type:varchar(64);comment:省"`
	City          string `gorm:"not null;type:varchar(64);comment:市"`
	Region        string `gorm:"not null;type:varchar(64);comment:区"`
	DetailAddress string `gorm:"not null;type:varchar(256);comment:详细地址"`
}

type AddressInfo struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	DefaultStatus int8   `json:"default_status"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Region        string `json:"region"`
	DetailAddress string `json:"detail_address"`
}

func (address Address) TableName() string {
	return "tb_receive_address"
}

func CreateAddress(ctx context.Context, db *gorm.DB, address *Address) error {
	result := db.WithContext(ctx).Create(address)
	return result.Error
}

func ExistDefaultAddress(ctx context.Context, tx *gorm.DB, userId int32) (Address, error) {
	var address Address
	err := tx.WithContext(ctx).Where(&Address{UserId: userId, DefaultStatus: AddressDefaultStatusYes}).First(&address).Error
	if err != nil {
		return address, err
	}
	return address, nil
}

func UpdateAddress(ctx context.Context, db *gorm.DB, addr Address) error {
	return db.WithContext(ctx).Save(&addr).Error
}

func GetAddressList(ctx context.Context, db *gorm.DB, userId int32) ([]AddressInfo, error) {
	var addressList []AddressInfo
	err := db.WithContext(ctx).Model(Address{}).Select("id, name, phone_number, default_status, province, city, region, detail_address").Where(Address{UserId: userId}).Find(&addressList).Error
	if err != nil {
		return addressList, err
	}
	return addressList, nil
}
