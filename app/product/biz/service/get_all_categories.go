package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	redisClient "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type GetAllCategoriesService struct {
	ctx context.Context
} // NewGetAllCategoriesService new GetAllCategoriesService
func NewGetAllCategoriesService(ctx context.Context) *GetAllCategoriesService {
	return &GetAllCategoriesService{ctx: ctx}
}

// Run create note info
func (s *GetAllCategoriesService) Run(req *product.CategoryListReq) (resp *product.CategoryListResp, err error) {
	//从redis获取分类列表
	key := model.CategoryKey()
	var categories []*product.Category
	err = getCategoryListByRedis(s.ctx, key, &categories)
	if err == nil {
		return &product.CategoryListResp{
			Categories: categories,
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
		}, nil
	} else {
		//获取分布式锁
		lockKey := model.CategoryLockKey()
		uuidString := uuid.New().String()
		lock, err := model.SetLock(s.ctx, redisClient.RedisClient, lockKey, uuidString)
		if err != nil {
			klog.CtxInfof(s.ctx, "获取分布式锁失败, err: %v", err)
			return nil, err
		}
		if lock {
			//如果获取到锁，则从mysql获取分类列表
			var categoriesList []*model.Category
			err := model.GetAllCategoryByDb(mysql.DB, s.ctx, &categoriesList)
			if err != nil {
				klog.CtxErrorf(s.ctx, "从mysql获取分类列表失败, err: %v", err)
				return nil, err
			}
			klog.CtxInfof(s.ctx, "从mysql获取分类列表成功, categories: %v", categoriesList)
			categoriesStr, err := sonic.MarshalString(categoriesList)
			if err != nil {
				klog.CtxErrorf(s.ctx, "序列化分类列表失败, err: %v", err)
				return nil, err
			}
			//再次判断锁是否存在
			luaScript := `
				local lockKey = KEYS[1]
				local key = KEYS[2]
				local uuid = ARGV[1]
				local value = ARGV[2]
				if redis.call("get", lockKey) == uuid then
					redis.call("set", key,value)
					redis.call("del", lockKey)
					return 1
				else
					return 0
				end
			`
			keys := []string{lockKey, key}
			args := []interface{}{uuidString, categoriesStr}
			result, err := redisClient.RedisClient.Eval(s.ctx, luaScript, keys, args).Result()
			if err != nil {
				klog.CtxErrorf(s.ctx, "redis lua脚本执行失败, err: %v", err)
				return nil, err
			}
			if result.(int64) == 1 {
				klog.CtxInfof(s.ctx, "redis lua脚本执行成功, categories: %v", categoriesList)
			} else {
				klog.CtxInfof(s.ctx, "redis lua脚本执行失败, categories: %v", categoriesList)
			}
		}
		//重试三次
		maxRetry := 3
		for i := 0; i < maxRetry; i++ {
			err := getCategoryListByRedis(s.ctx, key, &categories)
			if i == maxRetry-1 {
				if err != nil {
					klog.CtxErrorf(s.ctx, "超过最大重试次数:%v, err: %v", maxRetry, errors.WithStack(err))
					return nil, err
				}
			}
		}
	}

	return &product.CategoryListResp{
		Categories: categories,
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}

func getCategoryListByRedis(ctx context.Context, key string, categories *[]*product.Category) (err error) {
	result, err := redisClient.RedisClient.Get(ctx, key).Result()
	if err != nil {
		klog.CtxErrorf(ctx, "redis获取分类列表失败, err: %v", err)
		return err
	}
	//如果有，则直接返回
	if result != "" && result != "null" {
		err := sonic.UnmarshalString(result, &categories)
		if err != nil {
			klog.CtxErrorf(ctx, "反序列化分类列表失败, err: %v", err)
			return err
		}
		return nil
	}
	return errors.New("result is " + result)
}
