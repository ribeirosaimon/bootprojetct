package service

import (
	"context"
	"sync"
)

var (
	_onceShortener    sync.Once
	_shortenerService *shortenerService
)

type Shortener interface {
	ShortenerUrl(ctx context.Context)
}

type shortenerService struct {
}

func NewShortener() *shortenerService {
	_onceShortener.Do(func() {
		_shortenerService = &shortenerService{}
	})
	return _shortenerService
}

func (h *shortenerService) ShortenerUrl(ctx context.Context) {

	return
}
