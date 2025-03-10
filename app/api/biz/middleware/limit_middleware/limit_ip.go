package limit_middleware

import (
	"context"
	"douyin_mall/api/biz/dal/redis"
	"douyin_mall/api/conf"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func LimitIpMiddleware() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		clientIP := c.ClientIP()
		hlog.CtxInfof(ctx, "接收到请求, IP: %s, 访问接口", clientIP)
		key := "limit_ip:" + clientIP
		//初始化令牌漏桶
		capacity := conf.GetConf().LimitBucket.IPBucket.Capacity
		rate := conf.GetConf().LimitBucket.IPBucket.Rate
		_, err := redis.RedisClient.EvalSha(ctx, redis.LimitScriptSha, []string{key}, "init", capacity, rate).Result()
		if err != nil {
			hlog.CtxErrorf(ctx, "初始化令牌桶失败, IP: %s, 错误: %v", clientIP, err)
		}
		result, err := redis.RedisClient.EvalSha(ctx, redis.LimitScriptSha, []string{key}, "getTokens", 1).Result()
		if err != nil || result.(int64) != 1 {
			hlog.CtxErrorf(ctx, "获取令牌失败, IP: %s, 错误: %v", clientIP, err)
			hlog.CtxInfof(ctx, "IP: %s, 被流控", clientIP)
			c.AbortWithStatusJSON(429, utils.H{"code": 429, "message": "Too many requests"})
			return
		}

		c.Next(ctx)

	}
}
