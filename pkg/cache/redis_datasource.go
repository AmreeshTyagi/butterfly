package cache

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisDataSource is a interface
type RedisDataSource interface {
	Get(key string) string
	Set(key string, value interface{}, expiration time.Duration)
	Del(key ...string)
}

type redisDataSource struct {
	Cache *redis.Client
}

// NewRedisDataSource is a instance
func NewRedisDataSource(cache *redis.Client) RedisDataSource {
	return &redisDataSource{
		Cache: cache,
	}
}

func (r *redisDataSource) Get(key string) string {
	d, _ := r.Cache.Get(key).Result()
	return d
}

func (r *redisDataSource) Set(key string, value interface{}, expiration time.Duration) {
	r.Cache.Set(key, value, expiration)
}

func (r *redisDataSource) Del(key ...string) {
	r.Cache.Del(key...)
}
