package shorturl

import (
	"context"
)

type Repository interface {
	// Create метод создания shorturl
	Create(ctx context.Context, shorturl ShortUrl) (string, error)
	// FindOne поиск shorturl по идентификатору
	FindOne(ctx context.Context, id string) (string, error)
}
