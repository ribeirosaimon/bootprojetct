package service

import (
	"context"
	"github.com/ribeirosaimon/bootprojetct/internal/model/entity"
	"github.com/ribeirosaimon/bootprojetct/internal/repository"
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
	shortenerRepository repository.Shortener
}

func NewShortener() *shortenerService {
	_onceShortener.Do(func() {
		_shortenerService = &shortenerService{
			shortenerRepository: repository.NewShortener(),
		}
	})
	return _shortenerService
}

func (h *shortenerService) ShortenerUrl(ctx context.Context) {
	h.shortenerRepository.ShortenerUrl(ctx, entity.ShortenerUrl{})
	return
}
