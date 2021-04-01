package cache

import (
	"blog/config"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var menoryCache *MemoryCache

type MemoryCache struct {
	cache cache.Cache  // Only expose the methods we want to make available
	mu    sync.RWMutex // For increment / decrement prevent reads and writes
}

func NewCacheFile() *MemoryCache {
	if menoryCache == nil {
		menoryCache = &MemoryCache{
			cache: *cache.New(config.Cache.ExpiredTime, config.Cache.CleanUpTime),
			mu:    sync.RWMutex{},
		}
	}

	return menoryCache
}

func (memoryCache *MemoryCache) Set(key string, value interface{}, expires time.Duration) error {
	memoryCache.mu.Lock()
	defer memoryCache.mu.Unlock()
	memoryCache.cache.Set(key, value, expires)
	return nil
}

func (memoryCache *MemoryCache) Get(key string) (interface{}, bool) {
	memoryCache.mu.RLock()

	defer memoryCache.mu.RUnlock()

	value, found := memoryCache.cache.Get(key)
	return value, found
}
