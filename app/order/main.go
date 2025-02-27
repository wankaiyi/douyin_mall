package main

import (
	"douyin_mall/common/constant"
	hotKeyClient "douyin_mall/common/infra/hot_key_client"
	"douyin_mall/common/middleware"
	"douyin_mall/common/mtl"
	"douyin_mall/common/serversuite"
	"douyin_mall/common/utils/env"
	"douyin_mall/order/biz/dal"
	"douyin_mall/order/biz/dal/redis"
	"douyin_mall/order/biz/task"
	"douyin_mall/order/infra/kafka"
	"douyin_mall/order/infra/rpc"
	"douyin_mall/order/utils"
	"github.com/xxl-job/xxl-job-executor-go"
	"net"
	"os"
	"time"

	"douyin_mall/order/conf"
	"douyin_mall/order/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	os.Setenv("TZ", "Asia/Shanghai")
	time.Local, _ = time.LoadLocation("")
	serviceName := conf.GetConf().Kitex.Service
	mtl.InitLog(
		conf.GetConf().Kitex.LogFileName,
		conf.GetConf().Kitex.LogMaxSize,
		conf.GetConf().Kitex.LogMaxBackups,
		conf.GetConf().Kitex.LogMaxAge,
		conf.LogLevel(),
		conf.GetConf().Kafka.ClsKafka.Usser,
		conf.GetConf().Kafka.ClsKafka.Password,
		conf.GetConf().Kafka.ClsKafka.TopicId,
		serviceName,
	)
	mtl.InitTracing(serviceName, conf.GetConf().Jaeger.Endpoint)
	mtl.InitMetrics(serviceName, conf.GetConf().Kitex.MetricsPort)
	dal.Init()
	rpc.InitClient()
	kafka.Init()
	utils.InitSnowflake()

	//启动hotKeyClient
	go hotKeyClient.Start(redis.RedisClient, constant.OrderService)
	xxljobInit()

	opts := kitexInit()

	svr := orderservice.NewServer(new(OrderServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	var address string
	if currentEnv, err := env.GetString("env"); err == nil && currentEnv == "prod" {
		address = "0.0.0.0:8888"
	} else {
		address = conf.GetConf().Kitex.Address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))
	opts = append(opts, server.WithMiddleware(middleware.ServerInterceptor))
	opts = append(opts, server.WithMiddleware(middleware.BuildRecoverPanicMiddleware(conf.GetConf().Alert.FeishuWebhook)))
	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: conf.GetConf().Kitex.Service,
	}))

	return
}

func xxljobInit() {
	if currentEnv, err := env.GetString("env"); err == nil && currentEnv == "dev" {
		return
	}
	exec := xxl.NewExecutor(
		xxl.ServerAddr(conf.GetConf().XxlJob.XxlJobAddress+"/xxl-job-admin"),
		xxl.AccessToken(conf.GetConf().XxlJob.AccessToken),
		xxl.ExecutorIp(conf.GetConf().XxlJob.ExecutorIp),
		xxl.ExecutorPort("7777"),
		xxl.RegistryKey("douyin-mall-order-service"),
	)
	exec.Init()
	server.RegisterShutdownHook(func() {
		exec.Stop()
	})

	exec.RegTask("CleanNodeIDTask", task.CleanNodeIDTask)
	exec.RegTask("CancelOrderTask", task.CancelOrderTask)

	err := exec.Run()
	if err != nil {
		return
	}

}
