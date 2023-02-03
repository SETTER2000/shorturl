package usecase

import (
	"errors"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"log"
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

// ShortLink принимает короткий URL и возвращает длинный (GET /api/{key})
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

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(sh entity.Shorturl) (string, error) {
	key := scripts.UniqueString()
	log.Printf("FOOO post /::: %s", sh.URL)
	err := uc.repo.Put(key, sh.URL)
	if err == nil {
		return key, nil
	}

	return "", ErrBadRequest
}

// Shorten принимает длинный URL и возвращает короткий (POST /api/shorten)
func (uc *ShorturlUseCase) Shorten(sh entity.Shorturl) (string, error) {
	key := scripts.UniqueString()
	err := uc.repo.Put(key, sh.URL)
	if err == nil {
		return key, nil
	}

	return "", ErrBadRequest
}
