package redis

import "example.com/m/internal/infrastructure/env"

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisConfig() Config {
	return Config{
		// Redisのデフォルトポートは6379
		Addr:     env.GetString("REDIS_ADDR", "localhost:6379"),
		Password: env.GetString("REDIS_PASSWORD", ""), // デフォルトは空
		DB:       env.GetInt("REDIS_DB", 0),           // デフォルトのDB番号は0
	}
}
