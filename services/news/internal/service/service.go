package service

import (
	"errors"
	"fmt"
	"github.com/crypto-pulse/news/internal/integration/crypto_panic"
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/crypto-pulse/news/internal/model"
	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
)

func GetNews(c *gin.Context, rdb *redis.Client) ([]*model.News, error) {
	news, err := crypto_panic.FetchNews()
	if err != nil {
		return nil, err
	}

	var filtered []*model.News
	for _, post := range news {
		res, err := rdb.GetNews(c, fmt.Sprintf("%d", post.Id))
		if err != nil && !errors.Is(err, goredis.Nil) {
			return nil, err
		}

		if res != "" {
			continue
		}

		filtered = append(filtered, post)

		if err = rdb.SaveNews(c, fmt.Sprintf("%d", post.Id), post); err != nil {
			return nil, err
		}
	}

	return filtered, nil
}
