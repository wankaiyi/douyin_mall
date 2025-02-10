package dal

import (
	"douyin_mall/payment/biz/dal/alipay"
	"douyin_mall/payment/biz/dal/mysql"
	"douyin_mall/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	alipay.Init()
}
