package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func InitRedis(cfg RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("failed connect to redis: %w", err)
	}

	return rdb, nil
}
