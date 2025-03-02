package add

import (
	"context"
	"douyin_mall/product/biz/dal/redis"
	productModel "douyin_mall/product/biz/model"
	"douyin_mall/product/infra/kafka/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

func UpdateProductToRedis(ctx context.Context, product *model.UpdateProductSendToKafka) (err error) {
	key := "product:" + strconv.FormatInt(product.ID, 10)
	//4 调用redis的set方法将数据导入到redis缓存中
	err = productModel.PushToRedis(ctx, productModel.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		LockStock:   product.LockStock,
	}, redis.RedisClient, key)
	if err != nil {
		klog.CtxErrorf(ctx, "redis push product to redis err:%v", err)
		return err
	}
	return nil
}
