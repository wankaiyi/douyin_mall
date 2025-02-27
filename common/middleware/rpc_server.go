package middleware

import (
	"context"
	"douyin_mall/common/constant"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
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

		return next(ctx, req, resp)
	}
}
