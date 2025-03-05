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

type LockProductQuantityService struct {
	ctx context.Context
} // NewLockProductQuantityService new LockProductQuantityService
func NewLockProductQuantityService(ctx context.Context) *LockProductQuantityService {
	return &LockProductQuantityService{ctx: ctx}
}

// Run create note info
func (s *LockProductQuantityService) Run(req *product.ProductLockQuantityRequest) (resp *product.ProductLockQuantityResponse, err error) {
	originProducts := req.Products
	var ids = make([]int64, 0)
	var productQuantityMap = make(map[int64]int64)
	for _, pro := range originProducts {
		ids = append(ids, pro.Id)
		productQuantityMap[pro.Id] = pro.Quantity
	}
	//pipeline
	var cmdMap = make(map[int64]*redis.IntCmd)
	pipeline := redisClient.RedisClient.Pipeline()
	for _, id := range ids {
		productKey := model.BaseInfoKey(s.ctx, id)
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
		err := lockQuantity(s.ctx, productQuantityMap)
		if err != nil {
			return nil, err
		}
		return &product.ProductLockQuantityResponse{
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
			go pushToRedis(s.ctx, id, &wg, *hasError)
		}
		wg.Wait()

		//如果有异常，则睡眠500ms，再次尝试锁库存
		if *hasError {
			time.Sleep(500 * time.Millisecond)
		}
		err := lockQuantity(s.ctx, productQuantityMap)
		if err != nil {
			klog.CtxErrorf(s.ctx, "redis锁库存时执行失败 %v，keysMap:%v", errors.WithStack(err), productQuantityMap)
			return nil, err
		}
		return &product.ProductLockQuantityResponse{
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
		}, nil
	}
}

func lockQuantity(ctx context.Context, productQuantityMap map[int64]int64) error {
	luaScript := `
		local function process_keys(keys, quantities)
			for i, key in ipairs(keys) do
				redis.call('expire', key, 300)
				local quantity = tonumber(quantities[i])
				local stock = tonumber(redis.call('HGET', key, 'stock') or 0)
				local lock_stock = tonumber(redis.call('HGET', key, 'lock_stock') or 0)
				if stock - lock_stock < quantity then
					return 2
				end
			end
			for i, key in ipairs(keys) do
				local quantity = tonumber(quantities[i])
				redis.call('HINCRBY', key, 'lock_stock',quantity)
			end
			return 1
		end
		return process_keys(KEYS, ARGV)
		`
	keys := make([]string, 0)
	args := make([]interface{}, 0)
	for id, quantity := range productQuantityMap {
		productKey := model.StockKey(ctx, id)
		keys = append(keys, productKey)
		args = append(args, quantity)
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

func pushToRedis(ctx context.Context, id int64, wg *sync.WaitGroup, hasError bool) {
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
					Name:        pro.ProductName,
					Description: pro.ProductDescription,
					Picture:     pro.ProductPicture,
					Price:       pro.ProductPrice,
					Stock:       pro.ProductStock,
					Sale:        pro.ProductSale,
					PublicState: pro.ProductPublicState,
					LockStock:   pro.ProductLockStock,
					CategoryId:  pro.CategoryID,
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
