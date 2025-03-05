package redis

import "fmt"

const (
	UserInfoCacheKey      = "user:info:%d"
	UserAddressesCacheKey = "user:addresses:%d"
	UserAddressHitRateKey = "user:get_receive_address:hit_rate"
	UserInfoHitRateKey    = "user:get_info:hit_rate"
)

func GetUserKey(userId int32) string {
	return fmt.Sprintf(UserInfoCacheKey, userId)
}

func GetUserAddressKey(userId int32) string {
	return fmt.Sprintf(UserAddressesCacheKey, userId)
}

func GetUserAddressHitRateKey() string {
	return UserAddressHitRateKey
}

func GetUserInfoHitRateKey() string {
	return UserInfoHitRateKey
}
