package route

import (
	"github.com/crypto-pulse/news/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", handler.GetNewsHandler)
}
