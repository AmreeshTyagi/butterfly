package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Config is a model
type Config struct {
	Port     int
	Database int
	Host     string
	Pass     string
}

// RedisCache is a interface
type RedisCache interface {
	Client() *redis.Client
}

type redisCache struct {
	Config Config
}

// NewRedisCache is a instance
func NewRedisCache(config Config) RedisCache {
	return &redisCache{
		Config: config,
	}
}

func (r *redisCache) Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Config.Host, r.Config.Port),
		Password: r.Config.Pass,
		DB:       r.Config.Database,
	})
	return client
}
