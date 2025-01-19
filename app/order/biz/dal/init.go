package dal

import (
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
