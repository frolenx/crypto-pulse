package route

import (
	"github.com/crypto-pulse/news/internal/handler"
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, rdb *redis.Client) {
	router.GET("/", func(c *gin.Context) {
		handler.GetNewsHandler(c, rdb)
	})
}
