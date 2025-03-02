package job

import (
	"douyin_mall/product/conf"
	"douyin_mall/product/infra/xxl/task"
	"github.com/cloudwego/kitex/server"
	"github.com/xxl-job/xxl-job-executor-go"
)

func XxlJobInit() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(conf.GetConf().XxlJob.XxlJobAddress+"/xxl-job-admin"),
		xxl.AccessToken(conf.GetConf().XxlJob.AccessToken),
		xxl.ExecutorIp(conf.GetConf().XxlJob.ExecutorIp),
		xxl.ExecutorPort("7777"),
		xxl.RegistryKey("douyin-mall-product-service"),
	)
	exec.Init()
	server.RegisterShutdownHook(func() {
		exec.Stop()
	})

	exec.RegTask("RefreshElastic", task.RefreshElastic)
	exec.RegTask("SyncStockToDb", task.SyncStockToDb)

	err := exec.Run()
	if err != nil {
		return
	}

}
