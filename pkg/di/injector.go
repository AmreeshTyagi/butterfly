package di

import (
	"github.com/casbin/casbin"
	"github.com/go-redis/redis"
	"github.com/prongbang/butterfly/pkg/cache"
	"github.com/prongbang/butterfly/pkg/database"
)

// ProvideConfig is provide instance
func ProvideConfig() database.Config {
	return database.Config{
		DriverName:   "postgres",
		Host:         "localhost",
		Port:         5432,
		User:         "postgres",
		Pass:         "1e8c8d0a-011b-43f2-a3a6-7fdfb8ad619a",
		DatabaseName: "butterflydb",
		SSLMode:      "disable",
	}
}

// ProvideCasbinConfig is provide instance
func ProvideCasbinConfig() database.CasbinConfig {
	return database.CasbinConfig{
		DriverName:   "postgres",
		Host:         "localhost",
		Port:         5432,
		User:         "postgres",
		Pass:         "1e8c8d0a-011b-43f2-a3a6-7fdfb8ad619a",
		DatabaseName: "butterflydb",
		SSLMode:      "disable",
		PolicyFile:   "./config/auth_model.conf",
	}
}

// ProvideRedisConfig is provide instance
func ProvideRedisConfig() cache.Config {
	return cache.Config{
		Host:     "localhost",
		Port:     6379,
		Pass:     "e7a22218-34ac-41f5-952c-b35fed9213d4",
		Database: 0,
	}
}

// ProvideCoreDatabase is provide instance
func ProvideCoreDatabase() database.Database {
	return database.NewCoreDatabase(ProvideConfig())
}

// ProvideCasbinDatabase is provide instance
func ProvideCasbinDatabase() database.CasbinDatabase {
	return database.NewCasbinDatabase(ProvideCasbinConfig())
}

// ProvideEnforcer is provide instance
func ProvideEnforcer() *casbin.Enforcer {
	return ProvideCasbinDatabase().NewEnforcer()
}

// ProvideRedisCache is provide instance
func ProvideRedisCache() cache.RedisCache {
	return cache.NewRedisCache(ProvideRedisConfig())
}

// ProvideRedisClient is provide instance
func ProvideRedisClient() *redis.Client {
	return ProvideRedisCache().Client()
}
