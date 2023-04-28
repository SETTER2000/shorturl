package usecase

import (
	"context"
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
	in := entity.Shorturl{
		//Slug: "sl-1",
		//UserID: "1",
		//URL:    "http://xxxzzz.ru",
	}

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
	//shorturlMock := repoMock.NewMockShorturl(ctl)
	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)

	// Запрос к бд
	in := entity.Shorturl{
		Slug: "sl-1",
		//UserID: "1",
		//URL:    "http://xxxzzz.ru",
	}

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
	shorturlUseCaseMock.EXPECT().Get(ctx, &in).Return(&mockResp, nil).Times(1)

	UseCase := New(shorturlUseCaseMock)
	//sh, err := UseCase.ShortLink(ctx, &in)
	sh, err := UseCase.repo.Get(ctx, &in)
	require.NoError(t, err)
	require.Equal(t, &expected, sh)
}
func TestShorturlUseCase_Post(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	//shorturlMock := repoMock.NewMockShorturl(ctl)
	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)

	// Запрос к бд
	in := entity.Shorturl{
		Slug:   "sl-1",
		UserID: "1",
		URL:    "http://xxxzzz.ru",
	}
	//repoErr := errors.New("db is down")
	//mockResp := entity.Shorturl{
	//	Slug:   "sl-1",
	//	UserID: "1",
	//	URL:    "http://xxxzzz.ru",
	//}

	// База должна вернуть в ответ это
	//expected := entity.Shorturl{
	//	Slug:   "sl-1",
	//	UserID: "1",
	//	URL:    "http://xxxzzz.ru",
	//}
	// ожидаем, что вернётся
	shorturlUseCaseMock.EXPECT().Post(ctx, &in).Return(nil).Times(1)

	UseCase := New(shorturlUseCaseMock)
	//sh, err := UseCase.ShortLink(ctx, &in)
	err := UseCase.repo.Post(ctx, &in)
	require.NoError(t, err)
	//require.Equal(t, &expected, sh)
}

//func TestGetError(t *testing.T) {
//
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)
//
//	repoErr := errors.New("db is down")
//	ctx := context.Background()
//
//	// Запрос к бд
//	in := entity.Shorturl{
//		Slug: "sl-1",
//		//UserID: "1",
//		//URL:    "http://xxxzzz.ru",
//	}
//	shorturlUseCaseMock.EXPECT().Get(ctx, &in).Return(nil, repoErr).Times(1)
//
//	Usecase := New(shorturlUseCaseMock)
//	sh, err := Usecase.repo.Get(ctx, &in)
//	require.Error(t, err)
//	require.EqualError(
//		t,
//		fmt.Errorf("%s", repoErr.Error()),
//		err.Error(),
//	)
//	require.Nil(t, sh)
//}

//func TestShorturlUseCase_ShortLink(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	ctx := context.Background()
//	//shorturlMock := repoMock.NewMockShorturl(ctl)
//	shorturlUseCaseMock := repoMock.NewMockShorturlRepo(ctl)
//
//	// Запрос к бд
//	in := entity.Shorturl{
//		Slug: "sl-1",
//		//UserID: "1",
//		//URL:    "http://xxxzzz.ru",
//	}
//
//	mockResp := entity.Shorturl{
//		Slug:   "sl-1",
//		UserID: "1",
//		URL:    "http://xxxzzz.ru",
//	}
//
//	// База должна вернуть в ответ это
//	expected := entity.Shorturl{
//		Slug:   "sl-1",
//		UserID: "1",
//		URL:    "http://xxxzzz.ru",
//	}
//	// ожидаем, что вернётся
//	shorturlUseCaseMock.EXPECT().Get(ctx, &in).Return(&mockResp, nil).Times(1)
//
//	UseCase := New(shorturlUseCaseMock)
//	//sh, err := UseCase.ShortLink(ctx, &in)
//	sh, err := UseCase.ShortLink(ctx, &in)
//	require.NoError(t, err)
//	require.Equal(t, &expected, sh)
//}
