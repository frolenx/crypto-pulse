package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/crypto-pulse/news/internal/model"
	"github.com/redis/go-redis/v9"
	"os"
)

type Client struct {
	rdb *redis.Client
}

func NewClient() (*Client, error) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return nil, errors.New("REDIS_ADDR environment variable not set")
	}

	return &Client{
		rdb: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}, nil
}

func (c *Client) SaveNews(ctx context.Context, key string, value *model.News) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key, data, 0).Err()
}

func (c *Client) GetNews(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}
