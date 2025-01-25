package dal

import (
	"douyin_mall/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
}
