package usecase

import (
	"context"
	"errors"
	"github.com/SETTER2000/shorturl/internal/entity"
)

// ErrNotFound ошибка в случаи отсутствия данных
// ErrAlreadyExists ошибка в случаи если данные уже существуют
// ErrBadRequest ошибка в случаи неправильного формата запроса и т.п.
var (
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrBadRequest     = errors.New("bad request")
	ErrUserIDRequired = errors.New("user id required")
)

// ShorturlUseCase -.
type ShorturlUseCase struct {
	repo IShorturlRepo
}

// New -.
func New(r IShorturlRepo) *ShorturlUseCase {
	return &ShorturlUseCase{
		repo: r,
	}
}

// Post .
func (uc *ShorturlUseCase) Post(ctx context.Context, sh *entity.Shorturl) error {
	if err := uc.repo.Post(ctx, sh); err != nil {
		return err
	}
	return nil
}

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(ctx context.Context, sh *entity.Shorturl) (string, error) {
	err := uc.repo.Put(ctx, sh)
	if err != nil {
		return "", err
	}
	return sh.Slug, nil
}

// ShortLink принимает короткий URL и возвращает длинный (GET /api/{key})
func (uc *ShorturlUseCase) ShortLink(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	sh, err := uc.repo.Get(ctx, sh)
	if err == nil {
		return sh, nil
	}
	return nil, ErrBadRequest
}

// UserAllLink принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error) {
	//if len(u.UserID) < 1 {
	//	return nil, ErrUserIDRequired
	//}
	u, err := uc.repo.GetAll(ctx, u)
	if err == nil {
		return u, nil
	}
	return nil, ErrBadRequest
}

// AllLink принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) AllLink() (entity.CountURLs, error) {
	c, err := uc.repo.GetAllUrls()
	if err != nil {
		return 0, ErrBadRequest
	}

	return c, nil
}

// AllUsers принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) AllUsers() (entity.CountUsers, error) {
	c, err := uc.repo.GetAllUsers()
	if err != nil {
		return 0, ErrBadRequest
	}

	return c, nil
}

// UserDelLink принимает короткий URL и возвращает длинный (DELETE /api/user/urls)
func (uc *ShorturlUseCase) UserDelLink(ctx context.Context, u *entity.User) error {
	err := uc.repo.Delete(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

// SaveService сохраняет данные при выключении сервиса
func (uc *ShorturlUseCase) SaveService() error {
	err := uc.repo.Save()
	if err != nil {
		return err
	}
	return nil
}

// ReadService читает данные из хранилища при загрузки сервиса
func (uc *ShorturlUseCase) ReadService() error {
	err := uc.repo.Read()
	if err != nil {
		return err
	}
	return nil
}
