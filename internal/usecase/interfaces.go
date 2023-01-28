// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"github.com/SETTER2000/shorturl/internal/entity"
	"net/http"
)

type (
	// Shorturl -.
	Shorturl interface {
		LongLink(entity.Shorturl) (string, error)
		ShortLink(w http.ResponseWriter, r *http.Request) (string, error)
	}

	// ShorturlRepo -.
	ShorturlRepo interface {
		Get(key string) (string, error)
		Put(key, value string) error
	}
)
