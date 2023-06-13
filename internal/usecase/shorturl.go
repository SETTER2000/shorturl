package usecase

import (
	"context"
	"errors"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/app/er"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
)

// ShorturlUseCase -.
type ShorturlUseCase struct {
	repo IShorturlRepo
	cfg  *config.Config
}

// New -.
func New(r IShorturlRepo, cfg *config.Config) *ShorturlUseCase {
	return &ShorturlUseCase{
		repo: r,
		cfg:  cfg,
	}
}

// Post .
func (uc *ShorturlUseCase) Post(ctx context.Context, sh *entity.Shorturl) (*entity.ShorturlResponse, error) {
	resp := &entity.ShorturlResponse{}
	sh.Config = uc.cfg
	if sh.Slug == "" {
		sh.Slug = scripts.UniqueString()
	}

	err := uc.repo.Post(ctx, sh)
	if err != nil {
		if errors.Is(err, er.ErrAlreadyExists) {
			response, err := uc.repo.Get(ctx, sh)
			if err != nil {
				return nil, er.ErrBadRequest
			}

			url := scripts.GetHost(uc.cfg.HTTP, response.Slug)
			resp.URL = url
			return resp, er.ErrStatusConflict
		} else {
			return nil, er.ErrBadRequest
		}
	}

	resp.URL = scripts.GetHost(uc.cfg.HTTP, sh.Slug)
	return resp, nil
}

// LongLink принимает длинный URL и возвращает короткий (PUT /api)
func (uc *ShorturlUseCase) LongLink(ctx context.Context, sh *entity.Shorturl) (entity.Slug, error) {
	//sh.Config = uc.cfg
	//sh.Slug = scripts.UniqueString()
	//sh.UserID = entity.UserID(ctx.Value(uc.cfg.AccessTokenName).(string))
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
	return nil, er.ErrBadRequest
}

// UserAllLink принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error) {
	u, err := uc.repo.GetAll(ctx, u)
	if err == nil {
		return u, nil
	}
	return nil, er.ErrBadRequest
}

// AllLink принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) AllLink() (entity.CountURLs, error) {
	c, err := uc.repo.GetAllUrls()
	if err != nil {
		return 0, er.ErrBadRequest
	}

	return c, nil
}

// AllUsers принимает короткий URL и возвращает длинный (GET /user/urls)
func (uc *ShorturlUseCase) AllUsers() (entity.CountUsers, error) {
	c, err := uc.repo.GetAllUsers()
	if err != nil {
		return 0, er.ErrBadRequest
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
