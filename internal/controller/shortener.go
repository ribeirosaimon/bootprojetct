package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/bootprojetct/internal/service"
	"github.com/ribeirosaimon/bootprojetct/web/http"
	"log"
	"sync"
)

var (
	_onceShortener       sync.Once
	_shortenerController *shortenerController
)

type Shortener interface {
	GetShortenerUrl(ctx *gin.Context)
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

func (h *shortenerController) GetShortenerUrl(ctx *gin.Context) {
	type teste struct {
		Url string `json:"url"`
	}
	body, err := http.GetBody[teste](ctx)
	if err != nil {
		http.NewResponse(ctx).BadRequest("")
	}
	log.Printf("%s", body)
	h.shortenerService.ShortenerUrl(ctx)
	return
}
