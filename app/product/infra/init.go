package infra

import (
	"douyin_mall/product/infra/cache"
	"douyin_mall/product/infra/elastic"
	"douyin_mall/product/infra/kafka"
	"douyin_mall/product/infra/rpc"
	"douyin_mall/product/infra/xxl"
)

func Init() {
	elastic.InitClient()
	kafka.Init()
	xxl.Init()
	rpc.InitClient()
	cache.Init()
}
