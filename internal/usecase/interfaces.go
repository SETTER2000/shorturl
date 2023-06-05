// Package usecase - слой usecase, интерфейсы, реализует бизнес-логику приложения,
// каждая логическая группа в собственном файле.
package usecase

import (
	"context"
	"github.com/SETTER2000/shorturl/internal/usecase/repo"

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
	AllLink() (entity.CountURLs, error)
	AllUsers() (entity.CountUsers, error)
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
	GetAllUrls() (entity.CountURLs, error)
	GetAllUsers() (entity.CountUsers, error)
	Delete(context.Context, *entity.User) error
	Read() error
	Save() error
}

// строка ниже не несёт функциональной нагрузки
// её можно убрать без последствий для работы программы
// это отладочная строка
// в этой строке приведением типов проверяем,
// реализует ли структура *batchPostProvider интерфейс BatchPostProvider —
// если нет или если методы прописаны неверно,
// то компилятор выдаст на этой строке ошибку типизации
var _ IShorturlRepo = (*repo.InMemory)(nil)
var _ IShorturlRepo = (*repo.InFiles)(nil)
var _ IShorturlRepo = (*repo.InSQL)(nil)
