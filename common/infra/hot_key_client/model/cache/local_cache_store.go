package cache

import (
	"context"
	v "douyin_mall/common/infra/hot_key_client/model/value"
	"github.com/allegro/bigcache/v3"
	"github.com/bytedance/sonic"
	"log"

	"sync"
	"time"
)

var (
	// LocalStore 全局变量
	LocalStore *LocalCacheStore
)

func init() {
	var mapper sync.Map
	LocalStore = &LocalCacheStore{
		ConcurrentMap: mapper,
	}
}

// LocalCacheStore 本地缓存存储，缓存key对应的value_model
type LocalCacheStore struct {
	ConcurrentMap sync.Map
}

// PutInLocalCacheStore 往本地缓存中存入value,value为默认值(空),并设置过期时间
func (localCacheStore *LocalCacheStore) PutInLocalCacheStore(key string, value v.ValueModel) error {
	marshal, _ := sonic.Marshal(value)
	localCache, err := bigcache.New(context.Background(), bigcache.Config{
		Shards:             1024,
		MaxEntriesInWindow: 5000,
		LifeWindow:         time.Duration(value.Duration) * time.Millisecond,
		CleanWindow:        time.Duration(value.Duration) * time.Millisecond,
		Verbose:            true,
	})
	err = localCache.Set(key, marshal)
	if err != nil {
		return err
	}
	localCacheStore.ConcurrentMap.Store(key, localCache)
	return nil
}

// GetDefaultValue 从本地缓存中获取数据,如果没有则返回空的结果体，建议配合IsHot一起使用
func (localCacheStore *LocalCacheStore) GetDefaultValue(key string) (v.ValueModel, bool) {
	if localCache, ok := localCacheStore.ConcurrentMap.Load(key); ok {
		cache := localCache.(*bigcache.BigCache)
		value, err := cache.Get(key)
		if err != nil {
			return v.ValueModel{}, false
		}
		var valueModel v.ValueModel
		err = sonic.Unmarshal(value, &valueModel)
		return valueModel, true

	}
	return v.ValueModel{}, false
}

func (localCacheStore *LocalCacheStore) Remove(key string) {
	value, ok := localCacheStore.ConcurrentMap.Load(key)
	if ok {
		cache := value.(*bigcache.BigCache)
		err := cache.Delete(key)
		if err != nil {
			log.Printf("Delete err:%v ,key:%s", err, key)
			return
		}
	}
}

func IsHot(key string) (ok bool) {

	value, o := LocalStore.ConcurrentMap.Load(key)
	if !o {
		return false
	}
	cache := value.(*bigcache.BigCache)
	_, emptyError := cache.Get(key)
	if emptyError != nil {
		return false
	}
	return true

}
