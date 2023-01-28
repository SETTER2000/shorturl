package usecase

import (
	"errors"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/SETTER2000/shorturl/internal/entity"
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

// ShortLink принимает короткий URL и возвращает длинный
func (uc *ShorturlUseCase) ShortLink(w http.ResponseWriter, r *http.Request) (string, error) {
	shorturl := chi.URLParam(r, "key")
	if shorturl == "" {
		return "", ErrBadRequest
	}
	shorturl, err := uc.repo.Get(shorturl)
	if err == nil {
		return shorturl, nil
	}
	return "", ErrNotFound
}

// LongLink принимает длинный URL и возвращает короткий
func (uc *ShorturlUseCase) LongLink(sh entity.Shorturl) (string, error) {
	key := scripts.UniqueString()
	err := uc.repo.Put(key, sh.URL)
	if err == nil {
		return key, nil
	}

	return "", ErrBadRequest
}
