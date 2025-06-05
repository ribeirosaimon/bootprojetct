package entity

import (
	"encoding/base64"
	"github.com/google/uuid"
	"hash/fnv"
	"strings"
	"time"
)

type ShortenerUrl struct {
	ID         uuid.UUID `json:"id"`
	ShortCode  string    `json:"short_code"`
	Original   string    `json:"original"`
	Clicks     int64     `json:"clicks"`
	HashSize   int64     `json:"hash_size"`
	CreatedBy  string    `json:"created_by"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
	LastAccess time.Time `json:"last_access"`
}

func (s *ShortenerUrl) HashUrlShort(input string, size int) string {
	h := fnv.New64a()
	h.Write([]byte(input))
	sum := h.Sum(nil)
	// base64 pode ser mais curto que hex
	encoded := base64.URLEncoding.EncodeToString(sum)
	// remove caracteres que não são letras/números
	result := strings.TrimRight(encoded, "=")
	if size > 0 && size < len(result) {
		return result[:size]
	}
	return result
}
