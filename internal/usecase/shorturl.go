package usecase

import (
	"errors"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"

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
	sProduct  ShorturlRepoFilesProduct
	sConsumer ShorturlRepoFilesConsumer
}

// New -.
func New(sp ShorturlRepoFilesProduct, sc ShorturlRepoFilesConsumer, r ShorturlRepo) *ShorturlUseCase {
	return &ShorturlUseCase{
		sProduct:  sp,
		sConsumer: sc,
		repo:      r,
	}
}

func (uc *ShorturlUseCase) Shorten(sh *entity.Shorturl) (string, error) {
	sh.Slug = scripts.UniqueString()
	_, ok := os.LookupEnv("FILE_STORAGE_PATH")
	if !ok {
		err := uc.repo.Post(sh)
		if err == nil {
			return sh.Slug, nil
		}
	} else {
		err := uc.sProduct.Post(sh)
		if err == nil {
			return sh.Slug, nil
		}
	}

	return "", ErrBadRequest
}

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(sh *entity.Shorturl) (string, error) {
	sh.Slug = scripts.UniqueString()
	_, ok := os.LookupEnv("FILE_STORAGE_PATH")
	if !ok {
		err := uc.repo.Put(sh)
		if err == nil {
			return sh.Slug, nil
		}
	} else {
		err := uc.sProduct.Put(sh)
		if err == nil {
			return sh.Slug, nil
		}
	}

	return "", ErrBadRequest
}

// ShortLink принимает короткий URL и возвращает длинный (GET /api/{key})
func (uc *ShorturlUseCase) ShortLink(w http.ResponseWriter, r *http.Request) (string, error) {
	shorturl := chi.URLParam(r, "key")
	if shorturl == "" {
		return "", ErrBadRequest
	}
	_, ok := os.LookupEnv("FILE_STORAGE_PATH")
	if !ok {
		URL, err := uc.repo.Get(shorturl)
		if err == nil {
			return URL, nil
		}
	} else {
		sh, err := uc.sConsumer.Get(shorturl)
		if err == nil {
			return sh.URL, nil
		}
	}

	return "", ErrBadRequest
}
