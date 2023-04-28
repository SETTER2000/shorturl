package usecase

import (
	"context"
	"errors"
	"github.com/SETTER2000/shorturl/internal/entity"
	repoMock "github.com/SETTER2000/shorturl/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShorturlUseCase_ShortLink(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ctx := context.Background()
	shorturlMock := repoMock.NewMockShorturl(ctl)

	// Запрос к бд
	in := entity.Shorturl{}

	mockResp := entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}

	// База должна вернуть в ответ это
	expected := entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}
	// ожидаем, что вернётся
	shorturlMock.EXPECT().ShortLink(ctx, &in).Return(&mockResp, nil).Times(1)

	sh, err := shorturlMock.ShortLink(ctx, &in)
	require.NoError(t, err)
	require.Equal(t, &expected, sh)
}
func TestShorturlUseCase_Get(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ctx := context.Background()
	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)

	// Запрос к бд
	in := &entity.Shorturl{
		Slug: "sl-1",
	}

	mockResp := &entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}

	// База должна вернуть в ответ это
	expected := &entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}
	// ожидаем, что вернётся
	shorturlUseCaseMock.EXPECT().Get(ctx, in).Return(mockResp, nil).Times(1)

	UseCase := New(shorturlUseCaseMock)
	sh, err := UseCase.repo.Get(ctx, in)
	require.NoError(t, err)
	require.Equal(t, expected, sh)
}
func TestShorturlUseCase_Post(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ctx := context.Background()

	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)

	// Запрос к бд
	in := &entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}

	// ожидаем, что вернётся
	shorturlUseCaseMock.EXPECT().Post(ctx, in).Return(nil).Times(1)
	UseCase := New(shorturlUseCaseMock)
	err := UseCase.repo.Post(ctx, in)
	require.NoError(t, err)
}
func TestShorturlUseCase_PostError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)

	// Запрос к бд
	in := &entity.Shorturl{}
	repoErr := errors.New("db is down")

	// ожидаем, что вернётся
	shorturlUseCaseMock.EXPECT().Post(ctx, in).Return(repoErr).Times(1)
	UseCase := New(shorturlUseCaseMock)
	err := UseCase.repo.Post(ctx, in)
	require.Error(t, err)
}

func TestShortLink(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	shorturlMock := repoMock.NewMockShorturl(ctl)
	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)

	//repoErr := errors.New("db is down")
	// Запрос к бд
	in := &entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}

	mockResp := &entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}

	// База должна вернуть в ответ это
	expected := &entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}
	// ожидаем, что вернётся
	shorturlMock.EXPECT().ShortLink(ctx, in).Return(mockResp, nil).Times(1)
	sh, err := shorturlMock.ShortLink(ctx, in)
	require.NoError(t, err)
	require.Equal(t, expected, sh)

	// ожидаем, что вернётся
	shorturlUseCaseMock.EXPECT().Get(ctx, in).Return(mockResp, nil).Times(1)
	UseCase := New(shorturlUseCaseMock)
	sh, err = UseCase.repo.Get(ctx, in)
	require.NoError(t, err)
	require.Equal(t, expected, sh)
}
