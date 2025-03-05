package util

import (
	"context"
	"douyin_mall/common/constant"
	keyModel "douyin_mall/common/infra/hot_key_client/model/key"
	"douyin_mall/common/infra/hot_key_client/processor"
	"douyin_mall/product/infra/cache/model"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"time"
)

func SmartSet(ctx context.Context, isHotKey bool, key string, value interface{}, model *model.Cache, ttl time.Duration) (err error) {
	if isHotKey {
		marshalString, err := sonic.MarshalString(value)
		if err != nil {
			klog.CtxErrorf(ctx, "value %v 序列化失败, err: %v", value, err)
			return err
		}
		model.Set(key, marshalString, ttl)
	}
	return nil
}

func IsHotKey(ctx context.Context, key string) bool {
	keyConf := keyModel.KeyConf{
		ServiceName: constant.ProductService,
		Threshold:   1,
		Interval:    10,
		Duration:    10,
	}
	hotKeyModel := keyModel.NewHotKeyModelWithConfig(key, &keyConf)
	isHotKey := processor.IsHotKey(*hotKeyModel)
	return isHotKey
}

func SmartGet(ctx context.Context, key string, cache *model.Cache) (interface{}, bool) {
	return cache.Get(key)
}

func AddHit(ctx context.Context, key string, subKeys []string, client *redis.Client) {
	pipeline := client.Pipeline()
	for _, subKey := range subKeys {
		pipeline.HIncrBy(ctx, key, subKey, 1)
	}
	pipeline.Exec(ctx)
}
