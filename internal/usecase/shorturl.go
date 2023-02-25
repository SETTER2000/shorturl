package usecase

import (
	"context"
	"errors"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
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

func (uc *ShorturlUseCase) Shorten(ctx context.Context, sh *entity.Shorturl) (string, error) {
	sh.UserID = ctx.Value(sh.Cookie.AccessTokenName).(string)
	err := uc.repo.Post(ctx, sh)
	if err == nil {
		return sh.Slug, nil
	}
	return "", ErrBadRequest
}

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(ctx context.Context, sh *entity.Shorturl) (string, error) {
	sh.Slug = scripts.UniqueString()
	err := uc.repo.Put(ctx, sh)
	if err == nil {
		return sh.Slug, nil
	}
	return "", ErrBadRequest
}

// ShortLink принимает короткий URL и возвращает длинный (GET /api/{key})
func (uc *ShorturlUseCase) ShortLink(w http.ResponseWriter, r *http.Request) (*entity.Shorturl, error) {
	shorturl := chi.URLParam(r, "key")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if shorturl == "" {
		return nil, ErrBadRequest
	}
	sh, err := uc.repo.Get(ctx, shorturl)
	if err == nil {
		return sh, nil
	}
	return nil, ErrBadRequest
}

// UserAllLink принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error) {
	u, err := uc.repo.GetAll(ctx, u)
	if err == nil {
		return u, nil
	}
	return nil, ErrBadRequest
}
