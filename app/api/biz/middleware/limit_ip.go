package middleware

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

func LimitIpMiddleware() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		forwardedFor := string(c.GetHeader("X-Forwarded-For"))
		var clientIP string
		if forwardedFor != "" {
			ips := strings.Split(forwardedFor, ", ")
			clientIP = ips[0]
		} else {
			clientIP = c.ClientIP()
		}
		klog.CtxInfof(ctx, "接收到请求, IP: %s, 访问接口", clientIP)
		entry, blockError := sentinel.Entry("limit_ip", sentinel.WithTrafficType(base.Inbound), sentinel.WithResourceType(base.ResTypeWeb), sentinel.WithArgs(c.Host()))
		if blockError != nil {
			utils.LocalIP()
			hlog.CtxInfof(ctx, "IP: %s, 被流控", clientIP)
			c.AbortWithStatusJSON(429, utils.H{"code": 429, "message": "Too many requests"})
			return
		}

		defer entry.Exit()
		c.Next(ctx)

	}
}
