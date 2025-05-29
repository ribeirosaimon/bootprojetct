package controller

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	_onceHealth       sync.Once
	_healthController *healthController
)

type HealthController interface {
	GetHealth(ctx *gin.Context)
}

type healthController struct {
}

func NewHealthController() *healthController {
	_onceHealth.Do(func() {
		_healthController = &healthController{}
	})
	return _healthController
}

func (h *healthController) GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
	return
}
