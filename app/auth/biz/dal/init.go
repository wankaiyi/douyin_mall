package dal

import (
	"douyin_mall/auth/biz/dal/mysql"
	"douyin_mall/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
