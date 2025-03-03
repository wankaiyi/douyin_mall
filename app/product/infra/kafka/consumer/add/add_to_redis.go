package add

import (
	"context"
	"douyin_mall/product/biz/dal/redis"
	productModel "douyin_mall/product/biz/model"
	"douyin_mall/product/infra/kafka/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

func AddProductToRedis(ctx context.Context, product *model.AddProductSendToKafka) (err error) {
	key := "product:" + strconv.FormatInt(product.ID, 10)
	//4 调用redis的set方法将数据导入到redis缓存中
	err = productModel.PushToRedisBaseInfo(ctx, productModel.Product{
		Base: productModel.Base{
			ID: product.ID,
		},
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Picture:     product.Picture,
		Stock:       product.Stock,
		LockStock:   product.LockStock,
	}, redis.RedisClient, key)
	if err != nil {
		klog.CtxErrorf(ctx, "redis hset error: %v", err)
		return err
	}

	return nil
}
