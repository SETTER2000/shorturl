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
	repo      ShorturlRepo
	repoFiles ShorturlRepoFiles
}

// New -.
func New(r ShorturlRepo, rf ShorturlRepoFiles) *ShorturlUseCase {
	return &ShorturlUseCase{
		repo:      r,
		repoFiles: rf,
	}
}

// Shorten принимает длинный URL и возвращает короткий (POST /api/shorten)
func (uc *ShorturlUseCase) Shorten(sh *entity.Shorturl) (string, error) {
	sh.Slug = scripts.UniqueString()
	err := uc.repo.Post(sh)
	if err == nil {
		return sh.Slug, nil
	}

	return "", ErrBadRequest
}

// ShortLink принимает короткий URL и возвращает длинный (GET /api/{key})
func (uc *ShorturlUseCase) ShortLink(w http.ResponseWriter, r *http.Request) (string, error) {
	shorturl := chi.URLParam(r, "key")
	if shorturl == "" {
		return "", ErrBadRequest
	}
	sh, err := uc.repoFiles.Get(shorturl)
	if err != nil {
		return "", ErrNotFound
	}
	return sh.URL, nil
}

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(sh entity.Shorturl) (string, error) {
	key := scripts.UniqueString()
	log.Printf("LongLink post / %s", sh.URL)
	err := uc.repo.Put(key, sh.URL)
	if err == nil {
		return key, nil
	}

	return "", ErrBadRequest
}
