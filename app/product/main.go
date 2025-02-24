package main

import (
	"context"
	"douyin_mall/common/constant"
	hotKeyClient "douyin_mall/common/infra/hot_key_client"
	"douyin_mall/common/middleware"
	"douyin_mall/common/mtl"
	"douyin_mall/common/serversuite"
	"douyin_mall/common/utils/env"
	"douyin_mall/product/biz/dal"
	"douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/task"
	"douyin_mall/product/infra/elastic"
	"douyin_mall/product/infra/kafka"
	"douyin_mall/product/infra/rpc"
	"github.com/xxl-job/xxl-job-executor-go"
	"net"
	"os"
	"time"

	"douyin_mall/product/conf"
	"douyin_mall/product/kitex_gen/product/productcatalogservice"
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
	//启动hotKeyClient
	go hotKeyClient.Start(redis.RedisClient, constant.ProductService)

	elastic.InitClient()
	kafka.InitClient()
	opts := kitexInit()
	rpc.InitClient()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)
	//将任务注册到xxl-job中
	go xxljobInit()
	taskInit()
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
	xxljobAddr := conf.GetConf().XxlJob.XxlJobAddress
	exec := xxl.NewExecutor(
		xxl.ServerAddr(xxljobAddr+"/xxl-job-admin"),
		xxl.AccessToken(conf.GetConf().XxlJob.AccessToken), //请求令牌(默认为空)
		xxl.ExecutorIp(conf.GetConf().XxlJob.ExecutorIp),   //可自动获取
		xxl.ExecutorPort("7777"),                           //默认9999（非必填）
		xxl.RegistryKey("douyin-mall-product-service"),     //执行器名称
		xxl.SetLogger(&logger{}),                           //自定义日志
	)
	exec.Init()
	exec.Use(customMiddleware)
	//设置日志查看handler
	exec.LogHandler(customLogHandle)
	//注册任务handler
	exec.RegTask("task.RefreshElastic", task.RefreshElastic)

	klog.Fatal(exec.Run())
}

// 自定义日志处理器
func customLogHandle(req *xxl.LogReq) *xxl.LogRes {
	return &xxl.LogRes{Code: xxl.SuccessCode, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  "这个是自定义日志handler",
		IsEnd:       true,
	}}
}

// xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	klog.CtxInfof(context.Background(), format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	klog.CtxErrorf(context.Background(), format, a...)
}
func (l *logger) Debug(format string, a ...interface{}) {
	klog.CtxDebugf(context.Background(), format, a...)
}
func (l *logger) Warn(format string, a ...interface{}) {
	klog.CtxWarnf(context.Background(), format, a...)
}

// 自定义中间件
func customMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {
		klog.CtxInfof(context.Background(), "xxl-job 定时任务启动")
		res := tf(cxt, param)
		klog.CtxInfof(context.Background(), "xxl-job 定时任务结束")
		return res
	}
}

func taskInit() {
	task.ProduceIndicesInit()
}
