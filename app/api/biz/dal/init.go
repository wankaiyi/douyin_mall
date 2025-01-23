package dal

import (
	"douyin_mall/api/biz/dal/mysql"
	"douyin_mall/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
