package entity

import (
	"github.com/google/uuid"
	"time"
)

type ShortenerUrl struct {
	ID         uuid.UUID `json:"id"`
	ShortCode  string    `json:"short_code"`
	Original   string    `json:"original"`
	Clicks     int64     `json:"clicks"`
	CreatedBy  string    `json:"created_by"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
	LastAccess time.Time `json:"last_access"`
}
