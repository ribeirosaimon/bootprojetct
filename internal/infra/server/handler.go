package server

import "github.com/ribeirosaimon/bootprojetct/internal/controller"

func (b *bootProjectServer) startHandlers() {
	healthController := controller.NewHealth()
	shortenerController := controller.NewShortener()

	b.engine.GET("/health", healthController.GetHealth)
	b.engine.POST("/shortener", shortenerController.ShortenerUrl)
}
