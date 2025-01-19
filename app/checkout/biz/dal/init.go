package dal

import (
	"douyin_mall/checkout/biz/dal/mysql"
	"douyin_mall/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
