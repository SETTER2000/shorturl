// Package usecase - слой usecase, интерфейсы, реализует бизнес-логику приложения,
// каждая логическая группа в собственном файле.
package usecase

import (
	"context"

	"github.com/SETTER2000/shorturl/internal/entity"
)

// IShorturl - интерфейс обработчиков.
//
//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=IShorturl
type IShorturl interface {
	Post(context.Context, *entity.Shorturl) error
	LongLink(context.Context, *entity.Shorturl) (string, error)
	ShortLink(context.Context, *entity.Shorturl) (*entity.Shorturl, error)
	UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error)
	UserDelLink(ctx context.Context, u *entity.User) error
	ReadService() error
	SaveService() error
}

// IShorturlRepo - интерфейс DB.
//
//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=IShorturlRepo
type IShorturlRepo interface {
	Post(context.Context, *entity.Shorturl) error
	Put(context.Context, *entity.Shorturl) error
	Get(context.Context, *entity.Shorturl) (*entity.Shorturl, error)
	GetAll(context.Context, *entity.User) (*entity.User, error)
	Delete(context.Context, *entity.User) error
	Read() error
	Save() error
}
