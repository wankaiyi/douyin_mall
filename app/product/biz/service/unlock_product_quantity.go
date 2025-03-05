package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	redisClient "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type UnlockProductQuantityService struct {
	ctx context.Context
} // NewUnlockProductQuantityService new UnlockProductQuantityService
func NewUnlockProductQuantityService(ctx context.Context) *UnlockProductQuantityService {
	return &UnlockProductQuantityService{ctx: ctx}
}

// Run create note info
func (s *UnlockProductQuantityService) Run(req *product.ProductUnLockQuantityRequest) (resp *product.ProductUnLockQuantityResponse, err error) {
	ids := make([]int64, 0)
	quantities := make([]int64, 0)
	for _, pro := range req.Products {
		ids = append(ids, pro.ProductId)
		quantities = append(quantities, pro.Quantity)
	}
	//pipeline
	var cmdMap = make(map[int64]*redis.IntCmd)
	pipeline := redisClient.RedisClient.Pipeline()
	for _, id := range ids {
		productKey := model.StockKey(s.ctx, id)
		cmd := pipeline.Exists(s.ctx, productKey)
		cmdMap[id] = cmd
	}
	_, err = pipeline.Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	exist := true
	var notExistIds []int64 = make([]int64, 0)
	for id, cmd := range cmdMap {
		if cmd.Val() == 0 {
			notExistIds = append(notExistIds, id)
			exist = false
		}
	}

	if exist {
		err := unlockQuantity(s.ctx, ids, quantities)
		if err != nil {
			return nil, err
		}
		return &product.ProductUnLockQuantityResponse{
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
		}, nil
	} else {
		var hasError = new(bool)
		e := true
		hasError = &e
		var wg sync.WaitGroup
		for _, id := range notExistIds {
			wg.Add(1)
			go pushToRedisUnlock(s.ctx, id, &wg, *hasError)
		}
		wg.Wait()

		//如果有异常，则睡眠500ms，再次尝试锁库存
		if *hasError {
			time.Sleep(500 * time.Millisecond)
		}
		err := unlockQuantity(s.ctx, ids, quantities)
		if err != nil {
			klog.CtxErrorf(s.ctx, "redis锁库存时执行失败 %v", errors.WithStack(err))
			return nil, err
		}
		return &product.ProductUnLockQuantityResponse{
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
		}, nil
	}
}

func unlockQuantity(ctx context.Context, ids []int64, quantity []int64) error {
	keys := make([]string, 0)
	for _, id := range ids {
		keys = append(keys, model.StockKey(ctx, id))
	}
	luaScript := `
		local function process_keys(keys, quantities)
			for i, key in ipairs(keys) do
				if redis.call('EXISTS', key) == 0 then
					return 0
				else
					redis.call('EXPIRE', key, 600)
				end
			end
			for i, key in ipairs(keys) do
				local quantity = tonumber(quantities[i])
				redis.call('HINCRBY', key, 'lock_stock',-quantity)
			end
			return 1
		end
		return process_keys(KEYS, ARGV)
		`
	args := []interface{}{}
	for _, q := range quantity {
		args = append(args, q)
	}

	eval := redisClient.RedisClient.Eval(ctx, luaScript, keys, args)
	result, err := eval.Result()
	if err != nil {
		return err
	}
	if result.(int64) != 1 {
		return errors.New("锁库存失败")
	}
	return nil
}

func pushToRedisUnlock(ctx context.Context, id int64, wg *sync.WaitGroup, hasError bool) {
	defer wg.Done()
	//get lock
	//lock里面设置uuid，
	uuidString := uuid.New().String()
	lockKey := model.StockLockKey(ctx, id)
	nx := redisClient.RedisClient.SetNX(ctx, lockKey, uuidString, 2*time.Second) //1000ms
	if nx.Err() != nil {
		klog.CtxErrorf(ctx, "redis setnx error, reason: %v", nx.Err())
		hasError = true
	}
	stockKey := model.StockKey(ctx, id)
	//获取锁成功
	if nx.Val() {
		//再次确认这个key是否存在
		exist := redisClient.RedisClient.Exists(ctx, stockKey)
		if exist.Err() != nil {
			hasError = true
			klog.CtxErrorf(ctx, "redis exist error, reason: %v", errors.WithStack(exist.Err()))
		}
		//key exists, do nothing
		if exist.Val() == 1 {
			return
		} else {
			//先从数据库获取数据
			list, err := model.SelectProductList(mysql.DB, ctx, []int64{id})
			if err != nil {
				hasError = true
				return
			}
			for _, pro := range list {
				//然后推送到redis
				err := model.PushToRedisStock(ctx, model.Product{
					Base: model.Base{
						ID: pro.ProductId,
					},
					Name:          pro.ProductName,
					Description:   pro.ProductDescription,
					Picture:       pro.ProductPicture,
					Price:         pro.ProductPrice,
					Stock:         pro.ProductStock,
					Sale:          pro.ProductSale,
					PublishStatus: pro.ProductPublicState,
					LockStock:     pro.ProductLockStock,
					CategoryId:    pro.CategoryID,
				}, redisClient.RedisClient, stockKey)
				if err != nil {
					hasError = true
					klog.CtxErrorf(ctx, "redis push error, reason: %v, productId: %v", errors.WithStack(err), id)
					return
				}
			}
		}
		//释放锁
		unlockScript := `
						if redis.call("GET", KEYS[1]) == ARGV[1] then
							return redis.call("DEL", KEYS[1])
						else
							return 0
						end
					`
		keys := []string{lockKey}
		args := []interface{}{uuidString}
		unlockResult, err := redisClient.RedisClient.Eval(ctx, unlockScript, keys, args).Result()
		if err != nil {
			hasError = true
			return
		}
		//如果锁删除失败
		if unlockResult.(int64) != 1 {
			hasError = true
			return
		}
	} else {
		klog.CtxInfof(ctx, "redis setnx error, reason: %v", nx.Err())
		return
	}

}
