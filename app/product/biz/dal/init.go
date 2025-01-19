package dal

import (
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
