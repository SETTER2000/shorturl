// Code generated by mockery v2.20.2. DO NOT EDIT.

package usecase

import (
	context "context"

	entity "github.com/SETTER2000/shorturl/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockIShorturl is an autogenerated mock type for the IShorturl type
type MockIShorturl struct {
	mock.Mock
}

type MockIShorturl_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIShorturl) EXPECT() *MockIShorturl_Expecter {
	return &MockIShorturl_Expecter{mock: &_m.Mock}
}

// AllLink provides a mock function with given fields:
func (_m *MockIShorturl) AllLink() (entity.CountURLs, error) {
	ret := _m.Called()

	var r0 entity.CountURLs
	var r1 error
	if rf, ok := ret.Get(0).(func() (entity.CountURLs, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() entity.CountURLs); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entity.CountURLs)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIShorturl_AllLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllLink'
type MockIShorturl_AllLink_Call struct {
	*mock.Call
}

// AllLink is a helper method to define mock.On call
func (_e *MockIShorturl_Expecter) AllLink() *MockIShorturl_AllLink_Call {
	return &MockIShorturl_AllLink_Call{Call: _e.mock.On("AllLink")}
}

func (_c *MockIShorturl_AllLink_Call) Run(run func()) *MockIShorturl_AllLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturl_AllLink_Call) Return(_a0 entity.CountURLs, _a1 error) *MockIShorturl_AllLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturl_AllLink_Call) RunAndReturn(run func() (entity.CountURLs, error)) *MockIShorturl_AllLink_Call {
	_c.Call.Return(run)
	return _c
}

// LongLink provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturl) LongLink(_a0 context.Context, _a1 *entity.Shorturl) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Shorturl) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIShorturl_LongLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LongLink'
type MockIShorturl_LongLink_Call struct {
	*mock.Call
}

// LongLink is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.Shorturl
func (_e *MockIShorturl_Expecter) LongLink(_a0 interface{}, _a1 interface{}) *MockIShorturl_LongLink_Call {
	return &MockIShorturl_LongLink_Call{Call: _e.mock.On("LongLink", _a0, _a1)}
}

func (_c *MockIShorturl_LongLink_Call) Run(run func(_a0 context.Context, _a1 *entity.Shorturl)) *MockIShorturl_LongLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Shorturl))
	})
	return _c
}

func (_c *MockIShorturl_LongLink_Call) Return(_a0 string, _a1 error) *MockIShorturl_LongLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturl_LongLink_Call) RunAndReturn(run func(context.Context, *entity.Shorturl) (string, error)) *MockIShorturl_LongLink_Call {
	_c.Call.Return(run)
	return _c
}

// Post provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturl) Post(_a0 context.Context, _a1 *entity.Shorturl) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturl_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type MockIShorturl_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.Shorturl
func (_e *MockIShorturl_Expecter) Post(_a0 interface{}, _a1 interface{}) *MockIShorturl_Post_Call {
	return &MockIShorturl_Post_Call{Call: _e.mock.On("Post", _a0, _a1)}
}

func (_c *MockIShorturl_Post_Call) Run(run func(_a0 context.Context, _a1 *entity.Shorturl)) *MockIShorturl_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Shorturl))
	})
	return _c
}

func (_c *MockIShorturl_Post_Call) Return(_a0 error) *MockIShorturl_Post_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturl_Post_Call) RunAndReturn(run func(context.Context, *entity.Shorturl) error) *MockIShorturl_Post_Call {
	_c.Call.Return(run)
	return _c
}

// ReadService provides a mock function with given fields:
func (_m *MockIShorturl) ReadService() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturl_ReadService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadService'
type MockIShorturl_ReadService_Call struct {
	*mock.Call
}

// ReadService is a helper method to define mock.On call
func (_e *MockIShorturl_Expecter) ReadService() *MockIShorturl_ReadService_Call {
	return &MockIShorturl_ReadService_Call{Call: _e.mock.On("ReadService")}
}

func (_c *MockIShorturl_ReadService_Call) Run(run func()) *MockIShorturl_ReadService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturl_ReadService_Call) Return(_a0 error) *MockIShorturl_ReadService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturl_ReadService_Call) RunAndReturn(run func() error) *MockIShorturl_ReadService_Call {
	_c.Call.Return(run)
	return _c
}

// SaveService provides a mock function with given fields:
func (_m *MockIShorturl) SaveService() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturl_SaveService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveService'
type MockIShorturl_SaveService_Call struct {
	*mock.Call
}

// SaveService is a helper method to define mock.On call
func (_e *MockIShorturl_Expecter) SaveService() *MockIShorturl_SaveService_Call {
	return &MockIShorturl_SaveService_Call{Call: _e.mock.On("SaveService")}
}

func (_c *MockIShorturl_SaveService_Call) Run(run func()) *MockIShorturl_SaveService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturl_SaveService_Call) Return(_a0 error) *MockIShorturl_SaveService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturl_SaveService_Call) RunAndReturn(run func() error) *MockIShorturl_SaveService_Call {
	_c.Call.Return(run)
	return _c
}

// ShortLink provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturl) ShortLink(_a0 context.Context, _a1 *entity.Shorturl) (*entity.Shorturl, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Shorturl
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) (*entity.Shorturl, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) *entity.Shorturl); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Shorturl)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Shorturl) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIShorturl_ShortLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShortLink'
type MockIShorturl_ShortLink_Call struct {
	*mock.Call
}

// ShortLink is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.Shorturl
func (_e *MockIShorturl_Expecter) ShortLink(_a0 interface{}, _a1 interface{}) *MockIShorturl_ShortLink_Call {
	return &MockIShorturl_ShortLink_Call{Call: _e.mock.On("ShortLink", _a0, _a1)}
}

func (_c *MockIShorturl_ShortLink_Call) Run(run func(_a0 context.Context, _a1 *entity.Shorturl)) *MockIShorturl_ShortLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Shorturl))
	})
	return _c
}

func (_c *MockIShorturl_ShortLink_Call) Return(_a0 *entity.Shorturl, _a1 error) *MockIShorturl_ShortLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturl_ShortLink_Call) RunAndReturn(run func(context.Context, *entity.Shorturl) (*entity.Shorturl, error)) *MockIShorturl_ShortLink_Call {
	_c.Call.Return(run)
	return _c
}

// UserAllLink provides a mock function with given fields: ctx, u
func (_m *MockIShorturl) UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, u)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) (*entity.User, error)); ok {
		return rf(ctx, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) *entity.User); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIShorturl_UserAllLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserAllLink'
type MockIShorturl_UserAllLink_Call struct {
	*mock.Call
}

// UserAllLink is a helper method to define mock.On call
//   - ctx context.Context
//   - u *entity.User
func (_e *MockIShorturl_Expecter) UserAllLink(ctx interface{}, u interface{}) *MockIShorturl_UserAllLink_Call {
	return &MockIShorturl_UserAllLink_Call{Call: _e.mock.On("UserAllLink", ctx, u)}
}

func (_c *MockIShorturl_UserAllLink_Call) Run(run func(ctx context.Context, u *entity.User)) *MockIShorturl_UserAllLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User))
	})
	return _c
}

func (_c *MockIShorturl_UserAllLink_Call) Return(_a0 *entity.User, _a1 error) *MockIShorturl_UserAllLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturl_UserAllLink_Call) RunAndReturn(run func(context.Context, *entity.User) (*entity.User, error)) *MockIShorturl_UserAllLink_Call {
	_c.Call.Return(run)
	return _c
}

// UserDelLink provides a mock function with given fields: ctx, u
func (_m *MockIShorturl) UserDelLink(ctx context.Context, u *entity.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturl_UserDelLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserDelLink'
type MockIShorturl_UserDelLink_Call struct {
	*mock.Call
}

// UserDelLink is a helper method to define mock.On call
//   - ctx context.Context
//   - u *entity.User
func (_e *MockIShorturl_Expecter) UserDelLink(ctx interface{}, u interface{}) *MockIShorturl_UserDelLink_Call {
	return &MockIShorturl_UserDelLink_Call{Call: _e.mock.On("UserDelLink", ctx, u)}
}

func (_c *MockIShorturl_UserDelLink_Call) Run(run func(ctx context.Context, u *entity.User)) *MockIShorturl_UserDelLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User))
	})
	return _c
}

func (_c *MockIShorturl_UserDelLink_Call) Return(_a0 error) *MockIShorturl_UserDelLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturl_UserDelLink_Call) RunAndReturn(run func(context.Context, *entity.User) error) *MockIShorturl_UserDelLink_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockIShorturl interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIShorturl creates a new instance of MockIShorturl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIShorturl(t mockConstructorTestingTNewMockIShorturl) *MockIShorturl {
	mock := &MockIShorturl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
