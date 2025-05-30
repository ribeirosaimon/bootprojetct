package repository

import (
	"context"
	"github.com/ribeirosaimon/bootprojetct/internal/infra/database"
	"github.com/ribeirosaimon/bootprojetct/internal/model/entity"
	"sync"
)

var (
	_onceShortener       sync.Once
	_shortenerRepository *shortenerRepository
)

type Shortener interface {
	ShortenerUrl(ctx context.Context, url entity.ShortenerUrl) (*entity.ShortenerUrl, error)
}

type shortenerRepository struct {
	mysqlConnection database.BootProjectConnection
}

func NewShortener() *shortenerRepository {
	_onceShortener.Do(func() {
		_shortenerRepository = &shortenerRepository{
			mysqlConnection: database.NewMysqlConnection(),
		}
	})
	return _shortenerRepository
}

func (h *shortenerRepository) ShortenerUrl(ctx context.Context, url entity.ShortenerUrl) (*entity.ShortenerUrl, error) {

	// TODO como executar as queries
	// h.mysqlConnection.GetConnection().Query()
	return nil, nil
}
