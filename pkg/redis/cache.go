package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Set(ctx context.Context, key string, valuer string, expiration time.Duration) error {
	err := r.client.Set(ctx, key, valuer, expiration).Err()

	return err
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) || err != nil {
		return "", nil
	}

	return value, nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()

	return err
}
