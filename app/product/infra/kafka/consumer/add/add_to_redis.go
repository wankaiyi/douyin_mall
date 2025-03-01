package add

import (
	"context"
	"douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/infra/kafka/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

func AddProductToRedis(ctx context.Context, product *model.AddProductSendToKafka) (err error) {
	key := "product:" + strconv.FormatInt(product.ID, 10)
	//4 调用redis的set方法将数据导入到redis缓存中
	_, err = redis.RedisClient.HSet(ctx, key, map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"price":       product.Price,
		"picture":     product.Picture,
		"description": product.Description,
	}).Result()
	if err != nil {
		klog.CtxErrorf(ctx, "redis hset error: %v", err)
		return err
	}

	return nil
}
