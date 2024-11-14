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

func (r *Redis) Set(key string, valuer string, expiration time.Duration) error {
	ctx := context.Background()
	err := r.client.Set(ctx, key, valuer, expiration).Err()

	return err
}

func (r *Redis) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) || err != nil {
		return "", nil
	}

	return value, nil
}

func (r *Redis) Delete(key string) error {
	ctx := context.Background()
	err := r.client.Del(ctx, key).Err()

	return err
}
