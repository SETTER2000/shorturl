package repo

import (
	"context"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockShorturlRepo struct {
	mock.Mock
}

func (mock *MockShorturlRepo) Post(context.Context, *entity.Shorturl) error {
	args := mock.Called()
	return args.Error(1)
}
func (mock *MockShorturlRepo) Put(context.Context, *entity.Shorturl) error {
	args := mock.Called()
	return args.Error(1)
}
func (mock *MockShorturlRepo) Get(context.Context, *entity.Shorturl) (*entity.Shorturl, error) {
	args := mock.Called()
	result := args.Get(1)
	return result.(*entity.Shorturl), args.Error(1)
}
func (mock *MockShorturlRepo) GetAll(context.Context, *entity.User) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(1)
	return result.(*entity.User), args.Error(1)
}
func (mock *MockShorturlRepo) Delete(context.Context, *entity.User) error {
	args := mock.Called()
	//result := args.Get(0)
	return args.Error(1)
}
