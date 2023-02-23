package usecase

import (
	"errors"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrBadRequest    = errors.New("bad request")
)

// ShorturlUseCase -.
type ShorturlUseCase struct {
	repo ShorturlRepo
}

// New -.
func New(r ShorturlRepo) *ShorturlUseCase {
	return &ShorturlUseCase{
		repo: r,
	}
}

func (uc *ShorturlUseCase) Shorten(sh *entity.Shorturl) (string, error) {
	sh.Slug = scripts.UniqueString()
	err := uc.repo.Post(sh)
	if err == nil {
		return sh.Slug, nil
	}
	return "", ErrBadRequest
}

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(sh *entity.Shorturl) (string, error) {
	sh.Slug = scripts.UniqueString()
	err := uc.repo.Put(sh)
	if err == nil {
		return sh.Slug, nil
	}
	return "", ErrBadRequest
}

// ShortLink принимает короткий URL и возвращает длинный (GET /api/{key})
func (uc *ShorturlUseCase) ShortLink(w http.ResponseWriter, r *http.Request) (*entity.Shorturl, error) {
	shorturl := chi.URLParam(r, "key")
	if shorturl == "" {
		return nil, ErrBadRequest
	}
	sh, err := uc.repo.Get(shorturl)
	if err == nil {
		return sh, nil
	}
	return nil, ErrBadRequest
}

// UserAllLink принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) UserAllLink(u *entity.User) (*entity.User, error) {
	u, err := uc.repo.GetAll(u)
	if err == nil {
		return u, nil
	}
	return nil, ErrBadRequest
}
