package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/bootprojetct/internal/service"
	"sync"
)

var (
	_onceShortener       sync.Once
	_shortenerController *shortenerController
)

type Shortener interface {
	ShortenerUrl(ctx *gin.Context)
}

type shortenerController struct {
	shortenerService service.Shortener
}

func NewShortener() *shortenerController {
	_onceShortener.Do(func() {
		_shortenerController = &shortenerController{
			shortenerService: service.NewShortener(),
		}
	})
	return _shortenerController
}

func (h *shortenerController) ShortenerUrl(ctx *gin.Context) {
	//TODO
	h.shortenerService.ShortenerUrl(ctx)
	return
}
