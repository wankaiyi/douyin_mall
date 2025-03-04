// Code generated by hertz generator.

package main

import (
	"context"
	"douyin_mall/api/biz/middleware"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/api/mtl/log"
	"douyin_mall/common/mtl"
	"douyin_mall/common/utils/env"
	"douyin_mall/common/utils/feishu"
	"fmt"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"time"

	"douyin_mall/api/biz/router"
	"douyin_mall/api/conf"
	sentinelPlugin "github.com/alibaba/sentinel-golang/pkg/adapters/hertz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
)

var ServiceName string

func main() {
	os.Setenv("TZ", "Asia/Shanghai")
	time.Local, _ = time.LoadLocation("")

	ServiceName = conf.GetConf().Hertz.Service
	p := provider.NewOpenTelemetryProvider(
		provider.WithEnableMetrics(false),
		provider.WithServiceName(ServiceName),
		provider.WithExportEndpoint(conf.GetConf().Jaeger.Endpoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertztracing.NewServerTracer()

	registry, registryInfo := mtl.InitMetrics(ServiceName, conf.GetConf().Hertz.MetricsPort)
	defer registry.Deregister(registryInfo)

	rpc.InitClient()

	var address string
	if currentEnv, err := env.GetString("env"); err == nil && currentEnv == "prod" {
		address = "0.0.0.0:8888"
	} else {
		address = conf.GetConf().Hertz.Address
	}

	h := server.New(server.WithHostPorts(address), tracer)
	log.InitLog(ServiceName,
		conf.LogLevel(), conf.GetConf().Hertz.LogFileName, conf.GetConf().Hertz.LogMaxSize, conf.GetConf().Hertz.LogMaxBackups, conf.GetConf().Hertz.LogMaxAge,
		h,
		conf.GetConf().Kafka.ClsKafka.Usser, conf.GetConf().Kafka.ClsKafka.Password, conf.GetConf().Kafka.ClsKafka.TopicId)

	registerMiddleware(h, cfg)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	//hotkey 测试接口
	//要先启动hotkey worker,这里redis的地址要和worker的地址一致
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       0,
	//})
	//hotKeyClient.Start(client, constant.ApiServiceName)
	//h.GET("/isHot", func(c context.Context, ctx *app.RequestContext) {
	//	keyConf := key.NewKeyConf1(constant.ApiServiceName)
	//	hotkeyModel := key.NewHotKeyModelWithConfig("api", &keyConf)
	//	isHotKey := processor.IsHotKey(*hotkeyModel)
	//	if isHotKey {
	//		ctx.JSON(consts.StatusOK, utils.H{"isHotKey": "hot key"})
	//	} else {
	//		ctx.JSON(consts.StatusOK, utils.H{"isHotKey": "not hot key"})
	//	}
	//
	//})
	//h.GET("/set", func(c context.Context, ctx *app.RequestContext) {
	//	err := processor.SetDirectly("api", *value.NewValueModel(2000, "test set"))
	//	if err != nil {
	//		ctx.JSON(consts.StatusOK, utils.H{"setDirectly err": err.Error()})
	//	}
	//	ctx.JSON(consts.StatusOK, utils.H{"setDirectly": "success"})
	//})
	//
	//h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
	//	v := processor.Get("api")
	//	if v != nil {
	//		ctx.JSON(consts.StatusOK, utils.H{"api": v})
	//		return
	//	}
	//	ctx.JSON(consts.StatusOK, utils.H{"api": "not found"})
	//})
	//
	//h.GET("/remove", func(c context.Context, ctx *app.RequestContext) {
	//	processor.Remove("api")
	//	ctx.JSON(consts.StatusOK, utils.H{"api": "success"})
	//
	//})

	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz, cfg *hertztracing.Config) {

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(FeishuAlertRecoveryHandler)))

	// cores
	h.Use(cors.New(cors.Config{

		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{

			"*"},
		//准许使用的请求方式
		AllowMethods: []string{

			"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{

			"Origin", "access_token", "refresh_token", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{

			"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))

	h.Use(hertztracing.ServerMiddleware(cfg))
	h.Use(middleware.TraceLogMiddleware())
	h.Use(middleware.AuthorizationMiddleware())

	_, err := hotspot.LoadRules([]*hotspot.Rule{
		{
			Resource:        "limit_ip",
			MetricType:      hotspot.QPS,
			ControlBehavior: hotspot.Reject,
			Threshold:       1,
			ParamIndex:      0,
			DurationInSec:   1,
		},
	})
	if err != nil {
		hlog.Errorf("init sentinel error: %v", err)
	}

	h.Use(sentinelPlugin.SentinelServerMiddleware(sentinelPlugin.WithServerBlockFallback(func(ctx context.Context, a *app.RequestContext) {
		a.JSON(consts.StatusTooManyRequests, utils.H{"code": 429, "message": "Server is Busy"})
	})))
	h.Use(middleware.LimitIpMiddleware())
}

func FeishuAlertRecoveryHandler(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
	currentEnv, getEnvErr := env.GetString("env")
	if getEnvErr != nil {
		hlog.CtxErrorf(ctx, getEnvErr.Error())
	} else if currentEnv != "dev" {
		errMsg := fmt.Sprintf("接口%s发生错误：\nerr=%v\nstack=%s", c.Request.Path(), err, stack)
		hlog.CtxErrorf(ctx, "接口%s发生错误：\nerr=%v\nstack=%s", c.Request.Path(), err, stack)
		feishu.SendFeishuAlert(ctx, conf.GetConf().Alert.FeishuWebhook, errMsg)
	} else {
		// dev环境打印日志
		hlog.CtxErrorf(ctx, "接口%s发生错误：\nerr=%v\nstack=%s", c.Request.Path(), err, stack)
	}
	c.AbortWithStatusJSON(500, utils.H{"code": 500, "message": "Internal Server Error"})
}
