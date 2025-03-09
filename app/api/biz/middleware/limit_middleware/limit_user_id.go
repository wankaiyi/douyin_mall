package limit_middleware

import (
	"context"
	"douyin_mall/api/biz/dal/redis"
	"douyin_mall/api/conf"
	"douyin_mall/common/constant"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func LimitUserIdMiddleware() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		id := ctx.Value(constant.UserId)
		if id == nil {
			hlog.CtxInfof(ctx, "未登录")
			c.Next(ctx)
			return
		}
		userId := id.(int32)
		hlog.CtxInfof(ctx, "接收到请求, userId: %d, 访问接口", userId)
		key := fmt.Sprintf("limit_userId:%d", userId)
		//初始化令牌漏桶
		capacity := conf.GetConf().LimitBucket.UserIdBucket.Capacity
		rate := conf.GetConf().LimitBucket.UserIdBucket.Rate
		_, err := redis.RedisClient.EvalSha(ctx, redis.LimitScriptSha, []string{key}, "init", capacity, rate).Result()
		if err != nil {
			hlog.CtxErrorf(ctx, "初始化令牌桶失败, userId: %d, 错误: %v", userId, err)
		}
		result, err := redis.RedisClient.EvalSha(ctx, redis.LimitScriptSha, []string{key}, "getTokens", 1).Result()
		if err != nil || result.(int64) != 1 {
			hlog.CtxErrorf(ctx, "获取令牌失败, userId: %d, 错误: %v", userId, err)
			hlog.CtxInfof(ctx, "userId: %d, 被流控", userId)
			c.AbortWithStatusJSON(429, utils.H{"code": 429, "message": "Too many requests"})
			return
		}

		c.Next(ctx)

	}
}
