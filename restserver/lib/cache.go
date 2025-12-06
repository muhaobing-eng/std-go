package lib

import (
	"github.com/muhaobing-eng/std-go/go-common/cache"
	"github.com/muhaobing-eng/std-go/restserver/config"

	"github.com/go-redis/redis/v8"
)

var (
	cacheSingleInstance *redis.Client
)

func InitCache(cfg config.CacheConfig) error {
	cache, err := cache.New(cfg.GetCacheOption())
	if err != nil {
		return err
	}
	cacheSingleInstance = cache
	return nil
}

// GetRedis returns the redis instance
func GetRedis() *redis.Client {
	return cacheSingleInstance
}
