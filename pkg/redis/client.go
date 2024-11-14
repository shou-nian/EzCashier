package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/shou-nian/EzCashier/pkg/utils"
)

func OpenRedisConnection() (*Redis, error) {
	ctx := context.Background()

	// Build Redis connection URL.
	connURL, err := utils.ConnectionURLBuilder(utils.RedisConnection)
	if err != nil {
		return nil, err
	}

	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     connURL,
		Password: "", // no password set
		DB:       0,  // use DB0
	})

	// Check connection
	err = client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Redis{client: client}, nil
}
