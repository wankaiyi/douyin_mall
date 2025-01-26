package redis

import "fmt"

const (
	refreshTokenKey = "token:refresh:%d"
	accessTokenKey  = "token:access:%d"
)

func GetRefreshTokenKey(userId int32) string {
	return fmt.Sprintf(refreshTokenKey, userId)
}

func GetAccessTokenKey(userId int32) string {
	return fmt.Sprintf(accessTokenKey, userId)
}
