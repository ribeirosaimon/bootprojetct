package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/ribeirosaimon/bootprojetct/internal/infra/database"
	"github.com/ribeirosaimon/bootprojetct/internal/model/entity"
	"sync"
)

var (
	_onceShortener       sync.Once
	_shortenerRepository *shortenerRepository
)

type Shortener interface {
	GetAllShortenerUrl(ctx context.Context) ([]entity.ShortenerUrl, error)
	GetShortenerUrlById(ctx context.Context, id uuid.UUID) (*entity.ShortenerUrl, error)
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

func (h *shortenerRepository) GetAllShortenerUrl(ctx context.Context) ([]entity.ShortenerUrl, error) {
	query, err := h.mysqlConnection.GetConnection().Query(`
		SELECT
		    s.id,
		    s.short_code,
		    s.original,
		    s.clicks,
		    s.hash_size,
		    s.created_by,
		    s.is_active,
		    s.created_at,
		    s.expired_at,
		    s.last_access
		FROM shortener_url s
`)
	if err != nil {
		return nil, err
	}
	var shortenerUrls []entity.ShortenerUrl
	for query.Next() {
		var url entity.ShortenerUrl
		if err = query.Scan(
			&url.ID,
			&url.ShortCode,
			&url.Original,
			&url.Clicks,
			&url.HashSize,
			&url.CreatedBy,
			&url.IsActive,
			&url.CreatedAt,
			&url.ExpiredAt,
			&url.LastAccess,
		); err != nil {
			return nil, err
		}
		shortenerUrls = append(shortenerUrls, url)
	}
	return shortenerUrls, nil
}

func (h *shortenerRepository) GetShortenerUrlById(ctx context.Context, id uuid.UUID) (*entity.ShortenerUrl, error) {
	query := h.mysqlConnection.GetConnection().QueryRow(`
		SELECT
		    s.id,
		    s.short_code,
		    s.original,
		    s.clicks,
		    s.hash_size,
		    s.created_by,
		    s.is_active,
		    s.created_at,
		    s.expired_at,
		    s.last_access
		FROM shortener_url s WHERE s.id = ?
	`, id)
	var url entity.ShortenerUrl
	if err := query.Scan(
		&url.ID,
		&url.ShortCode,
		&url.Original,
		&url.Clicks,
		&url.HashSize,
		&url.CreatedBy,
		&url.IsActive,
		&url.CreatedAt,
		&url.ExpiredAt,
		&url.LastAccess,
	); err != nil {
		return nil, err
	}
	return &url, nil
}
