package middleware

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func LimitIpMiddleware() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {
		entry, blockError := sentinel.Entry("limit_ip", sentinel.WithTrafficType(base.Inbound), sentinel.WithResourceType(base.ResTypeWeb), sentinel.WithArgs(ctx.ClientIP()))
		if blockError != nil {
			hlog.CtxInfof(c, "IP: %s, 被流控", ctx.ClientIP())
			ctx.AbortWithStatusJSON(429, utils.H{"code": 429, "message": "Too many requests"})
			return
		}

		defer entry.Exit()
		ctx.Next(c)

	}
}
