package middleware

import (
	"context"
	"douyin_mall/common/constant"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"time"
)

func ServerInterceptor(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		// traceId是一定有的
		traceId, ok := metainfo.GetPersistentValue(ctx, constant.TraceId)
		if ok {
			context.WithValue(ctx, constant.TraceId, traceId)
		} /*else {
			panic("未找到traceId")
		}*/
		userId, ok := metainfo.GetPersistentValue(ctx, constant.UserId)
		if ok {
			context.WithValue(ctx, constant.UserId, userId)
		}

		ri := rpcinfo.GetRPCInfo(ctx)
		if ri == nil {
			return next(ctx, req, resp)
		}

		start := time.Now()
		klog.Infof("RPC请求目标服务: %s, 方法: %s, 请求参数: %+v",
			ri.To().ServiceName(), ri.To().Method(), req)

		err := next(ctx, req, resp)

		duration := time.Since(start).Milliseconds()
		if err != nil {
			klog.Errorf("RPC请求目标服务Failed: %s, 方法: %s, 耗时: %v毫秒, 错误信息: %v",
				ri.To().ServiceName(), ri.To().Method(), duration, err)
		} else {
			klog.Infof("RPC请求目标服务成功响应: %s, 方法:: %s, 耗时: %v毫秒, 响应: %+v",
				ri.To().ServiceName(), ri.To().Method(), duration, resp)
		}

		return err
	}
}
