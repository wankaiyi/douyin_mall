package model

import (
	"sync"
	"time"
)

// CacheItem 代表缓存中的一个项
type CacheItem struct {
	value      interface{}
	expiration int64 // 过期时间，Unix 时间戳
}

// Cache 是一个本地缓存类
type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

// NewCache 创建一个新的缓存实例
func NewCache() *Cache {
	cache := &Cache{
		items: make(map[string]CacheItem),
	}
	go cache.cleanupExpiredItems()
	return cache
}

// Set 添加或更新缓存项
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(ttl).UnixNano()
	c.items[key] = CacheItem{
		value:      value,
		expiration: expiration,
	}
}

// Get 获取缓存项
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	// 检查是否过期
	if time.Now().UnixNano() > item.expiration {
		delete(c.items, key)
		return nil, false
	}

	return item.value, true
}

// Delete 删除缓存项
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// cleanupExpiredItems 定期清理过期的缓存项
func (c *Cache) cleanupExpiredItems() {
	for {
		time.Sleep(time.Minute) // 每分钟检查一次

		c.mu.Lock()
		now := time.Now().UnixNano()
		for key, item := range c.items {
			if now > item.expiration {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
