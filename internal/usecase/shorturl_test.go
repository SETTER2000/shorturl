package usecase

import (
	"context"
	"github.com/SETTER2000/shorturl/internal/entity"
	repoMock "github.com/SETTER2000/shorturl/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInSQL_Get(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	//log := logger.New("Debug")
	//shMock := repoMock.NewMockShorturl(ctl)
	repo := repoMock.NewMockShorturlRepo(ctl)

	// Запрос к бд
	req := entity.Shorturl{
		Slug: "sl123",
		URL:  "http://xxxzzz.ru",
	}
	// База должна вернуть в ответ это
	res := entity.Shorturl{
		Slug: "sl",
		URL:  "http://xxxzzz.ru",
	}
	// ожидаем, что вернётся
	repo.EXPECT().Get(ctx, &req).Return(&res, nil).Times(1)

	UseCase := New(repo)
	sh, err := UseCase.repo.Get(ctx, &req)
	require.NoError(t, err)
	_ = sh
}
