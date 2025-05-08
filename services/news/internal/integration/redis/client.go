package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"os"
)

type Client struct {
	redisDb *redis.Client
}

func NewClient() (*Client, error) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return nil, errors.New("REDIS_ADDR environment variable not set")
	}

	return &Client{
		redisDb: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}, nil
}

func (c *Client) Save(ctx context.Context, key string, data []byte) error {
	return c.redisDb.Set(ctx, key, data, 0).Err()
}

func (c *Client) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := c.redisDb.Exists(ctx, key).Result()
	return exists == 1, err
}
