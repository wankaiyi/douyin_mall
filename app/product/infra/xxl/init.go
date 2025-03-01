package xxl

import (
	"douyin_mall/product/conf"
	"douyin_mall/product/infra/xxl/job"
)

func Init() {
	if conf.GetConf().Env != "dev" {
		job.XxlJobInit()
	}
}
