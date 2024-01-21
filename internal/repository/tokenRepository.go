package repository

import (
	"context"
	"errors"
	"gameReview/internal/config"
	"gameReview/internal/storage/redis"
	redis2 "github.com/redis/go-redis/v9"
)

var (
	ErrNotFound = errors.New("not found")
)

type TokenRepository struct {
	rdb *redis.RedisStorage
	cfg *config.Config
}

func NewTokenRepository(redis *redis.RedisStorage, cfg *config.Config) *TokenRepository {
	return &TokenRepository{rdb: redis, cfg: cfg}
}

func (tr TokenRepository) Token(ctx context.Context, email string) (string, error) {
	token, err := tr.rdb.Get(ctx, email).Result()
	if err != nil {
		if errors.Is(err, redis2.Nil) {
			return "", ErrNotFound
		} else {
			return "", err
		}
	}
	return token, nil
}

func (tr TokenRepository) Create(ctx context.Context, email, token string) (string, error) {
	token, err := tr.rdb.Set(ctx, email, token, tr.cfg.TokenTTL).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}
