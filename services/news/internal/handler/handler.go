package handler

import (
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/crypto-pulse/news/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNewsHandler(c *gin.Context, rdb *redis.Client) {
	news, err := service.GetNews(c, rdb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}
