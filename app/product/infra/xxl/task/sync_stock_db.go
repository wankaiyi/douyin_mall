package task

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	redisClient "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/xxl-job/xxl-job-executor-go"
	"strconv"
	"sync"
)

func SyncStockToDb(cxt context.Context, param *xxl.RunReq) string {
	//扫描redis所有库存的key，将数据全部拿下来
	var wg = sync.WaitGroup{}
	for {
		result, cursor, err := redisClient.RedisClient.Scan(cxt, 0, model.StockPattern+"*", 100).Result()
		if err != nil {
			klog.CtxErrorf(cxt, "RedisClient.Scan() error: %v", err)
			return err.Error()
		}
		for _, key := range result {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := syncToMysql(cxt, key)
				if err != nil {
					klog.CtxErrorf(cxt, "syncToMysql() error: %v", err)
				}
			}()
		}
		if cursor == 0 {
			break
		}
	}
	wg.Wait()
	//携程同步至数据库
	return "done"
}

func syncToMysql(ctx context.Context, key string) (err error) {
	klog.CtxInfof(ctx, "syncToMysql key:%v", key)
	id, err := strconv.Atoi(key[len(model.StockPattern):])
	if err != nil {
		klog.CtxInfof(ctx, "字符串转数字失败, key:%v,err:%v", key, err)
		return err
	}
	uuidString := uuid.New().String()
	lockKey := model.StockLockKey(ctx, int32(id))
	lock, err := model.SetLock(ctx, redisClient.RedisClient, lockKey, uuidString)
	if err != nil {
		return err
	}
	if lock {
		klog.CtxInfof(ctx, "id:%v上锁成功", id)
		result, err := redisClient.RedisClient.HGetAll(ctx, model.StockKey(ctx, int32(id))).Result()
		if err != nil {
			klog.CtxInfof(ctx, "redisClient.HGetAll() error: %v", err)
			return err
		}
		stock, err := strconv.Atoi(result["stock"])
		if err != nil {
			klog.CtxInfof(ctx, "stock数字失败,stock:%v,err: %v", result["stock"], err)
			return err
		}
		lockStock, err := strconv.Atoi(result["lock_stock"])
		if err != nil {
			klog.CtxInfof(ctx, "lock_stock数字失败,stock:%v,err: %v", result["lock_stock"], err)
			return err
		}
		err = model.UpdateProduct(mysql.DB, ctx, &model.Product{
			Base:      model.Base{ID: int32(id)},
			Stock:     int32(stock),
			LockStock: int32(lockStock),
		})
		if err != nil {
			klog.CtxInfof(ctx, "将数据更新到数据库失败,err:%v", err)
			return err
		}
		err = model.SafeDeleteLock(ctx, redisClient.RedisClient, lockKey, uuidString)
		if err != nil {
			klog.CtxInfof(ctx, "删除锁失败,err:%v", err)
		}
	} else {
		klog.CtxInfof(ctx, "id:%v上锁失败", id)
	}
	return nil
}
