package main

import (
	"context"
	"douyin_mall/cart/biz/dal"
	"douyin_mall/common/infra/nacos"
	"douyin_mall/common/utils/env"
	"douyin_mall/common/utils/feishu"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"net"
	"time"

	"douyin_mall/cart/conf"
	"douyin_mall/cart/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	dal.Init()
	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

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
	opts = append(opts, server.WithMiddleware(buildCoreMiddleware(&server.Options{})))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	r := nacos.RegisterService()
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

func buildCoreMiddleware(opt *server.Options) endpoint.Middleware {
	return RecoverPanic
}

func RecoverPanic(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		err := next(ctx, req, resp)
		ri := rpcinfo.GetRPCInfo(ctx)
		endpointInfo := ri.To()
		if err != nil {
			currentEnv, getEnvErr := env.GetString("env")
			if getEnvErr != nil {
				klog.Error(getEnvErr.Error())
			} else if currentEnv != "dev" {
				feishuWebhook := conf.GetConf().Alert.FeishuWebhook
				errMsg := fmt.Sprintf("服务**%s**接口**%s**发生异常，错误信息：%+v", endpointInfo.ServiceName(), endpointInfo.Method(), err)
				feishu.SendFeishuAlert(ctx, feishuWebhook, errMsg)
			}
		}
		return err
	}
}
