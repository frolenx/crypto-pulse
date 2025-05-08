package route

import (
	"github.com/crypto-pulse/news/internal/handler"
	"github.com/crypto-pulse/news/internal/integration/crypto_panic"
	"github.com/crypto-pulse/news/internal/integration/kafka/producer"
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, cryptoPanicApi *crypto_panic.Client, redisDb *redis.Client, kafkaProducer *producer.Producer) {
	newsHandler := handler.NewNewsHandler(cryptoPanicApi, redisDb, kafkaProducer)
	router.GET("/", newsHandler.ProcessLatestNews)
}
