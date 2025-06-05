package entity

import (
	"fmt"
	"log"
	"testing"
)

func BenchmarkShortenerUrl_HashUrlShort(b *testing.B) {
	s := &ShortenerUrl{}
	hashes := make(map[string]string) // shortCode => original

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Gere URLs diferentes para cada iteração
		original := fmt.Sprintf("https://minhaurl.com/artigo/%d", i)
		hash := s.HashUrlShort(original, 8)
		log.Printf("Tentativa %d com hash %s", i, hash)

		// Cheque colisão
		if exist, ok := hashes[hash]; ok && exist != original {
			b.Fatalf("Collision detected! SHA1: %s, urls: %s, %s", hash, exist, original)
		}

		hashes[hash] = original
	}
}
