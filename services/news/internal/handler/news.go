package handler

import (
	"github.com/crypto-pulse/news/internal/integration/crypto_panic"
	"github.com/crypto-pulse/news/internal/integration/kafka/producer"
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/crypto-pulse/news/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsHandler struct {
	newsService *service.NewsService
}

func NewNewsHandler(cryptoPanicApi *crypto_panic.Client, redisDb *redis.Client, kafkaProducer *producer.Producer) *NewsHandler {
	return &NewsHandler{
		newsService: service.NewNewsService(cryptoPanicApi, redisDb, kafkaProducer),
	}
}

func (h *NewsHandler) ProcessLatestNews(c *gin.Context) {
	err := h.newsService.FetchAndPublishNews(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
