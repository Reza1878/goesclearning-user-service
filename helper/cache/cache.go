package cache

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Reza1878/goesclearning/user-service/helper/fault"
	"github.com/redis/go-redis/v9"
)

func Set(ctx context.Context, client *redis.Client, key string, value interface{}, ttl time.Duration) error {
	err := client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fault.Custom(
			http.StatusUnprocessableEntity,
			fault.ErrUnprocessable,
			fmt.Sprintf("failed to save to Redis [key=%s]: %v", key, err))
	}
	return nil
}

func Exist(ctx context.Context, client *redis.Client, key string) (bool, error) {
	count, err := client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func Get(ctx context.Context, client *redis.Client, key string) (interface{}, error) {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}
