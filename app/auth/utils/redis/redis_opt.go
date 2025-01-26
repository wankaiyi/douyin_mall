package redis

import (
	"context"
	"douyin_mall/auth/biz/dal/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func SetVal(ctx context.Context, key string, val interface{}, expiration time.Duration) (string, error) {
	result, err := redis.RedisClient.Set(ctx, key, val, expiration).Result()
	if err != nil {
		klog.Error("redis set失败，", err)
	}
	return result, err
}

func GetVal(ctx context.Context, key string) (string, error) {
	result, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		klog.Error("redis get失败，", err)
	}
	return result, err
}
