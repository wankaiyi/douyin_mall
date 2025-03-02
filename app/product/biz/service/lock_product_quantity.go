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
	"strconv"
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
		cmd := pipeline.Exists(s.ctx, "product:"+strconv.FormatInt(id, 10))
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
		hasError := false
		var wg sync.WaitGroup
		for _, id := range notExistIds {
			wg.Add(1)
			go func(context context.Context, id int64, wg *sync.WaitGroup) {
				defer wg.Done()
				//get lock
				//lock里面设置uuid，
				uuidString := uuid.New().String()
				lockKey := "product:lock:" + strconv.FormatInt(id, 10)
				nx := redisClient.RedisClient.SetNX(s.ctx, lockKey, uuidString, 1*time.Millisecond) //1000ms
				if nx.Err() != nil {
					klog.CtxErrorf(s.ctx, "redis setnx error, reason: %v", nx.Err())
					hasError = true
				}
				productKey := "product:" + strconv.FormatInt(id, 10)
				//获取锁成功
				if nx.Val() {
					//再次确认这个key是否存在
					exist := redisClient.RedisClient.Exists(s.ctx, productKey)
					if exist.Err() != nil {
						hasError = true
						klog.CtxErrorf(s.ctx, "redis exist error, reason: %v", errors.WithStack(exist.Err()))
					}
					//key exists, do nothing
					if exist.Val() == 1 {
						return
					} else {
						//先从数据库获取数据
						list, err := model.SelectProductList(mysql.DB, s.ctx, []int64{id})
						if err != nil {
							hasError = true
							return
						}
						for _, pro := range list {
							//然后推送到redis
							err := model.PushToRedis(s.ctx, model.Product{
								ID:          pro.ProductId,
								Name:        pro.ProductName,
								Description: pro.ProductDescription,
								Picture:     pro.ProductPicture,
								Price:       pro.ProductPrice,
								Stock:       pro.ProductStock,
								Sale:        pro.ProductSale,
								PublicState: pro.ProductPublicState,
								LockStock:   pro.ProductLockStock,
								CategoryId:  pro.CategoryID,
								RealStock:   pro.RealStock,
							}, redisClient.RedisClient, "product:"+strconv.FormatInt(id, 10))
							if err != nil {
								hasError = true
								klog.CtxErrorf(s.ctx, "redis push error, reason: %v, productId: %v", errors.WithStack(err), id)
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
					unlockResult, err := redisClient.RedisClient.Eval(s.ctx, unlockScript, keys, args).Result()
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
					klog.CtxInfof(s.ctx, "redis setnx error, reason: %v", nx.Err())
					return
				}

			}(s.ctx, id, &wg)
		}
		wg.Wait()

		//如果有异常，则睡眠500ms，再次尝试锁库存
		if hasError {
			time.Sleep(500 * time.Millisecond)
		}
		err := lockQuantity(s.ctx, productQuantityMap)
		if err != nil {
			return nil, err
		}
		return &product.ProductLockQuantityResponse{
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
		}, nil
	}

	//productList, err := model.SelectProductList(mysql.DB, context.Background(), ids)
	////确定当前库存是否足够
	//canLock := true
	//var lowStockProductId int64
	//
	//for _, pro := range productList {
	//	//如果真实库存小于下单的数量，则库存锁定失败
	//	if pro.RealStock < productQuantityMap[pro.ProductId] {
	//		canLock = false
	//		lowStockProductId = pro.ProductId
	//		break
	//	}
	//}
	////如果库存锁定失败，则返回失败信息
	//if !canLock {
	//	klog.CtxInfof(s.ctx, "商品库存不足，无法锁定库存，productId：%v, quantity：%v", lowStockProductId, productQuantityMap[lowStockProductId])
	//	return &product.ProductLockQuantityResponse{
	//		StatusCode: 6014,
	//		StatusMsg:  constant.GetMsg(6014),
	//	}, nil
	//}
	////如果库存锁定成功，则更新库存信息
	//err = model.UpdateLockStock(mysql.DB, context.Background(), productQuantityMap)
	//if err != nil {
	//	klog.CtxErrorf(s.ctx, "更新库存失败，原因：%v", err)
	//	return nil, err
	//}
	//return &product.ProductLockQuantityResponse{
	//	StatusCode: 0,
	//	StatusMsg:  constant.GetMsg(0),
	//}, nil
}

func lockQuantity(ctx context.Context, productQuantityMap map[int64]int64) error {
	luaScript := `
		local function process_keys(keys, quantities)
			for i, key in ipairs(keys) do
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
		productKey := "product:" + strconv.FormatInt(id, 10)
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
