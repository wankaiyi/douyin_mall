package cache

import (
	"github.com/allegro/bigcache/v3"
	"sync"
)

var (
	concurrentMap = sync.Map{}
	Default       = "default"
)

func GetAllKeyCache(serviceName string) *bigcache.BigCache {

	if serviceName == "" {
		if _, ok := concurrentMap.Load(Default); !ok {
			allKeyCache := NewAllKeyCache()
			concurrentMap.Store(Default, allKeyCache)
		}
		value, _ := concurrentMap.Load(Default)
		return value.(*bigcache.BigCache)
	}
	if _, ok := concurrentMap.Load(serviceName); !ok {
		allKeyCache := NewAllKeyCache()
		concurrentMap.Store(serviceName, allKeyCache)
	}
	value, _ := concurrentMap.Load(serviceName)

	return value.(*bigcache.BigCache)
}
