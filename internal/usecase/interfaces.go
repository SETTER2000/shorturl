// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/SETTER2000/shorturl/internal/entity"
	"net/http"
)

type (
	// Shorturl -.
	Shorturl interface {
		Shorten(context.Context, *entity.Shorturl) (string, error)
		LongLink(context.Context, *entity.Shorturl) (string, error)
		ShortLink(w http.ResponseWriter, r *http.Request) (*entity.Shorturl, error)
		UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error)
	}

	// ShorturlRepo -.
	ShorturlRepo interface {
		Post(context.Context, *entity.Shorturl) error
		Put(context.Context, *entity.Shorturl) error
		Get(ctx context.Context, key string) (*entity.Shorturl, error)
		GetAll(context.Context, *entity.User) (*entity.User, error)
	}
)
