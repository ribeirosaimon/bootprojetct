package entity

import (
	"github.com/google/uuid"
	"time"
)

type ShortenerUrl struct {
	ID         uuid.UUID
	ShortCode  string
	Original   string
	Clicks     int64
	CreatedBy  string
	IsActive   bool
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LastAccess time.Time
}
