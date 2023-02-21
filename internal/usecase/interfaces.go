// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"github.com/SETTER2000/shorturl/internal/entity"
	"net/http"
)

type (
	// Shorturl -.
	Shorturl interface {
		Shorten(*entity.Shorturl) (string, error)
		LongLink(*entity.Shorturl) (string, error)
		ShortLink(w http.ResponseWriter, r *http.Request) (*entity.Shorturl, error)
		UserAllLink(u *entity.User) (*entity.User, error)
	}

	// ShorturlRepo -.
	ShorturlRepo interface {
		Post(*entity.Shorturl) error
		Put(*entity.Shorturl) error
		Get(key string) (*entity.Shorturl, error)
		GetAll(*entity.User) (*entity.User, error)
	}

	Store interface {
		Set(key string, value []byte) error
		Get(key string) ([]byte, error)
		Delete(key string) error
	}
)
