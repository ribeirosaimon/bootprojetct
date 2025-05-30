package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/bootprojetct/internal/controller"
)

func StartHandlers(e *gin.Engine) {
	healthController := controller.NewHealth()
	shortenerController := controller.NewShortener()

	e.GET("/health", healthController.GetHealth)
	e.POST("/shortener", shortenerController.ShortenerUrl)
}
