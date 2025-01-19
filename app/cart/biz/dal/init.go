package dal

import (
	"douyin_mall/cart/biz/dal/mysql"
	"douyin_mall/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
