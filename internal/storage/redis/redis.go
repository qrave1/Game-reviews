package redis

import (
	"fmt"
	"gameReview/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	*redis.Client
}

func NewRedisStorage(cfg *config.Config) *RedisStorage {
	c := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", cfg.Redis.Host),
		Password: "",
		DB:       0,
	})

	return &RedisStorage{Client: c}
}
