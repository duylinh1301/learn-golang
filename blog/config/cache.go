package config

import "time"

var (
	Cache = NewConfigCache()
)

type configCache struct {
	ExpiredTime time.Duration
	CleanUpTime time.Duration
}

func NewConfigCache() *configCache {
	return &configCache{
		ExpiredTime: 5 * time.Minute,
		CleanUpTime: 30 * time.Second,
	}
}
