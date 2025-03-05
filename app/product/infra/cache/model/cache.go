package model

import (
	"github.com/bytedance/sonic"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
)

// Cache 是一个基于 bigcache 的本地缓存类
type Cache struct {
	cache *bigcache.BigCache
	mu    sync.RWMutex
}

// NewCache 创建一个新的缓存实例
func NewCache() *Cache {
	// 配置 bigcache
	config := bigcache.Config{
		Shards:             1024,             // 分片数量
		LifeWindow:         10 * time.Minute, // 缓存项的生命周期
		CleanWindow:        1 * time.Minute,  // 清理过期缓存项的间隔
		MaxEntriesInWindow: 1000 * 10 * 60,   // 生命周期内的最大条目数
		MaxEntrySize:       500,              // 单个缓存项的最大大小（字节）
		HardMaxCacheSize:   128,              // 缓存的最大内存占用（MB）
		Verbose:            true,             // 是否打印日志
	}

	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err) // 初始化失败时直接 panic
	}

	return &Cache{
		cache: cache,
	}
}

// Set 添加或更新缓存项
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 将 value 转换为字节切片
	v, err := sonic.Marshal(value)
	if err != nil {
		return
	} // 假设 value 是结构体类型
	valueBytes, ok := value.([]byte)
	if !ok {
		valueBytes = v // 假设 value 是字符串类型
	}

	// 设置缓存项
	err = c.cache.Set(key, valueBytes)
	if err != nil {
		panic(err)
	}
}

// Get 获取缓存项
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// 获取缓存项
	valueBytes, err := c.cache.Get(key)
	if err != nil {
		if err == bigcache.ErrEntryNotFound {
			return nil, false // 缓存项不存在
		}
	}
	var value interface{}
	sonic.Unmarshal(valueBytes, &value) // 将字节切片转换为结构体类型

	// 返回缓存值
	return value, true
}

// Delete 删除缓存项
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 删除缓存项
	err := c.cache.Delete(key)
	if err != nil {
		panic(err)
	}
}
