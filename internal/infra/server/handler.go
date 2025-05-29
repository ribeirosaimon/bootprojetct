package server

import "github.com/ribeirosaimon/bootprojetct/internal/controller"

func (b *bootProjectServer) startHandlers() {
	healthController := controller.NewHealthController()

	b.engine.GET("/health", healthController.GetHealth)
}
