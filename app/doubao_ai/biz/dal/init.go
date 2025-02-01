package dal

import (
	"douyin_mall/doubao_ai/biz/dal/mysql"
	"douyin_mall/doubao_ai/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
