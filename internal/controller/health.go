package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/bootprojetct/web/http"
	"sync"
)

var (
	_onceHealth       sync.Once
	_healthController *healthController
)

type Health interface {
	GetHealth(ctx *gin.Context)
}

type healthController struct {
}

func NewHealth() *healthController {
	_onceHealth.Do(func() {
		_healthController = &healthController{}
	})
	return _healthController
}

func (h *healthController) GetHealth(ctx *gin.Context) {
	http.NewResponse(ctx).Success(gin.H{
		"status": "ok",
	})
	return
}
