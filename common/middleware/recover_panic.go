package middleware

import (
	"context"
	"douyin_mall/common/utils/env"
	"douyin_mall/common/utils/feishu"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func BuildRecoverPanicMiddleware(feishuWebhook string) endpoint.Middleware {
	return RecoverPanic(feishuWebhook)
}

func RecoverPanic(feishuWebhook string) func(next endpoint.Endpoint) endpoint.Endpoint {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) error {
			err := next(ctx, req, resp)
			ri := rpcinfo.GetRPCInfo(ctx)
			endpointInfo := ri.To()
			if err != nil {
				currentEnv, getEnvErr := env.GetString("env")
				if getEnvErr != nil {
					klog.CtxErrorf(ctx, getEnvErr.Error())
				} else if currentEnv != "dev" {
					errMsg := fmt.Sprintf("服务**%s**接口**%s**发生异常，错误信息：%+v", endpointInfo.ServiceName(), endpointInfo.Method(), err)
					klog.CtxErrorf(ctx, errMsg)
					feishu.SendFeishuAlert(ctx, feishuWebhook, errMsg)
				} else {
					klog.CtxErrorf(ctx, err.Error())
				}
			}
			return err
		}
	}
}
