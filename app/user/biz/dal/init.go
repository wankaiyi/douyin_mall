package dal

import (
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
