package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crypto-pulse/news/internal/integration/crypto_panic"
	"github.com/crypto-pulse/news/internal/integration/kafka/producer"
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
)

type NewsService struct {
	cryptoPanicApi *crypto_panic.Client
	redisDb        *redis.Client
	kafkaProducer  *producer.Producer
}

func NewNewsService(cryptoPanicApi *crypto_panic.Client, redisDb *redis.Client, kafkaProducer *producer.Producer) *NewsService {
	return &NewsService{
		cryptoPanicApi: cryptoPanicApi,
		redisDb:        redisDb,
		kafkaProducer:  kafkaProducer,
	}
}

func (s *NewsService) FetchAndPublishNews(c *gin.Context) error {
	news, err := s.cryptoPanicApi.FetchNews()
	if err != nil {
		return err
	}

	for _, post := range news {
		exists, err := s.redisDb.Exists(c, fmt.Sprintf("%d", post.Id))
		if err != nil && !errors.Is(err, goredis.Nil) {
			return err
		}

		if exists {
			continue
		}

		data, err := json.Marshal(post)
		if err != nil {
			return err
		}

		if err = s.redisDb.Save(c, fmt.Sprintf("%d", post.Id), data); err != nil {
			return err
		}

		s.kafkaProducer.Publish(fmt.Sprintf("%d", post.Id), data)
	}

	return nil
}
