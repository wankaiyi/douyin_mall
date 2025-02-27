package add

import (
	"context"
	"douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/kafka/model"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

func AddProductToRedis(ctx context.Context, product *model.AddProductSendToKafka) (err error) {
	pro := vo.ProductRedisDataVo{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Picture:     product.Picture,
	}
	key := "product:" + strconv.FormatInt(product.ID, 10)
	//4 调用redis的set方法将数据导入到redis缓存中
	marshal, err := sonic.MarshalString(pro)
	if err != nil {
		klog.Error("序列化失败", err)
		return err
	}
	_, err = redis.RedisClient.Set(ctx, key, marshal, 1*time.Hour).Result()
	if err != nil {
		klog.Error("redis set error", err)
		return err
	}

	return nil
}
