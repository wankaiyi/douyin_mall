package main

import (
	"douyin_mall/common/middleware"
	"douyin_mall/common/mtl"
	"douyin_mall/common/serversuite"
	"douyin_mall/common/utils/env"
	"douyin_mall/doubao_ai/biz/dal"
	"douyin_mall/doubao_ai/conf"
	"douyin_mall/doubao_ai/kitex_gen/doubao_ai/doubaoaiservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
	"os"
	"time"
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

	opts := kitexInit()

	svr := doubaoaiservice.NewServer(new(DoubaoAiServiceImpl), opts...)

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
