package middleware

import (
	"context"
	"douyin_mall/common/constant"
	utils "douyin_mall/common/utils/context"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func ClientInterceptor(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		ctx = utils.EnsurePersistentValue(ctx, constant.TraceId, true)
		ctx = utils.EnsurePersistentValue(ctx, constant.UserId, false)

		return next(ctx, req, resp)
	}
}
