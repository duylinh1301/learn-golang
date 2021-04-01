package cache

import (
	"blog/config"
	"fmt"

	"github.com/patrickmn/go-cache"
)

type CacheFile struct {
}

func NewCacheFile() *CacheFile {
	return &CacheFile{}
}

func Set() {
	c := cache.New(config.Cache.ExpiredTime, config.Cache.CleanUpTime)
	fmt.Println(c)
}
