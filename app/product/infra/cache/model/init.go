package model

var (
	LocalCache *Cache
)

func CacheInit() {
	LocalCache = NewCache()
}
