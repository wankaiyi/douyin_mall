package elastic

import (
	"douyin_mall/product/infra/elastic/check"
	"douyin_mall/product/infra/elastic/client"
)

func InitClient() {
	client.InitClient()
	check.ProduceIndicesInit()
}
