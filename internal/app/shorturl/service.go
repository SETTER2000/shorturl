package shorturl

import (
	"context"
)

// Service
// здесь не делаем интерфейс так как он нужен в основном для разной реализации
// например для бд, чтоб использовать разные базы данных
type Service struct {
	storage Repository
}

func (s *Service) Create(ctx context.Context, dto ShortUrl) (u string, err error) {
	return
}

func (s *Service) FindOne(ctx context.Context, id string) (u ShortUrl, err error) {
	return
}
