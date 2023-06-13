// Code generated by mockery v2.20.2. DO NOT EDIT.

package usecase

import (
	context "context"

	entity "github.com/SETTER2000/shorturl/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockIShorturlRepo is an autogenerated mock type for the IShorturlRepo type
type MockIShorturlRepo struct {
	mock.Mock
}

type MockIShorturlRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIShorturlRepo) EXPECT() *MockIShorturlRepo_Expecter {
	return &MockIShorturlRepo_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturlRepo) Delete(_a0 context.Context, _a1 *entity.User) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturlRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockIShorturlRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.User
func (_e *MockIShorturlRepo_Expecter) Delete(_a0 interface{}, _a1 interface{}) *MockIShorturlRepo_Delete_Call {
	return &MockIShorturlRepo_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *MockIShorturlRepo_Delete_Call) Run(run func(_a0 context.Context, _a1 *entity.User)) *MockIShorturlRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User))
	})
	return _c
}

func (_c *MockIShorturlRepo_Delete_Call) Return(_a0 error) *MockIShorturlRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturlRepo_Delete_Call) RunAndReturn(run func(context.Context, *entity.User) error) *MockIShorturlRepo_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturlRepo) Get(_a0 context.Context, _a1 *entity.Shorturl) (*entity.Shorturl, error) {
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

// MockIShorturlRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockIShorturlRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.Shorturl
func (_e *MockIShorturlRepo_Expecter) Get(_a0 interface{}, _a1 interface{}) *MockIShorturlRepo_Get_Call {
	return &MockIShorturlRepo_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *MockIShorturlRepo_Get_Call) Run(run func(_a0 context.Context, _a1 *entity.Shorturl)) *MockIShorturlRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Shorturl))
	})
	return _c
}

func (_c *MockIShorturlRepo_Get_Call) Return(_a0 *entity.Shorturl, _a1 error) *MockIShorturlRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturlRepo_Get_Call) RunAndReturn(run func(context.Context, *entity.Shorturl) (*entity.Shorturl, error)) *MockIShorturlRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturlRepo) GetAll(_a0 context.Context, _a1 *entity.User) (*entity.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) (*entity.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) *entity.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIShorturlRepo_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockIShorturlRepo_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.User
func (_e *MockIShorturlRepo_Expecter) GetAll(_a0 interface{}, _a1 interface{}) *MockIShorturlRepo_GetAll_Call {
	return &MockIShorturlRepo_GetAll_Call{Call: _e.mock.On("GetAll", _a0, _a1)}
}

func (_c *MockIShorturlRepo_GetAll_Call) Run(run func(_a0 context.Context, _a1 *entity.User)) *MockIShorturlRepo_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.User))
	})
	return _c
}

func (_c *MockIShorturlRepo_GetAll_Call) Return(_a0 *entity.User, _a1 error) *MockIShorturlRepo_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturlRepo_GetAll_Call) RunAndReturn(run func(context.Context, *entity.User) (*entity.User, error)) *MockIShorturlRepo_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUrls provides a mock function with given fields:
func (_m *MockIShorturlRepo) GetAllUrls() (entity.CountURLs, error) {
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

// MockIShorturlRepo_GetAllUrls_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUrls'
type MockIShorturlRepo_GetAllUrls_Call struct {
	*mock.Call
}

// GetAllUrls is a helper method to define mock.On call
func (_e *MockIShorturlRepo_Expecter) GetAllUrls() *MockIShorturlRepo_GetAllUrls_Call {
	return &MockIShorturlRepo_GetAllUrls_Call{Call: _e.mock.On("GetAllUrls")}
}

func (_c *MockIShorturlRepo_GetAllUrls_Call) Run(run func()) *MockIShorturlRepo_GetAllUrls_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturlRepo_GetAllUrls_Call) Return(_a0 entity.CountURLs, _a1 error) *MockIShorturlRepo_GetAllUrls_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturlRepo_GetAllUrls_Call) RunAndReturn(run func() (entity.CountURLs, error)) *MockIShorturlRepo_GetAllUrls_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUsers provides a mock function with given fields:
func (_m *MockIShorturlRepo) GetAllUsers() (entity.CountUsers, error) {
	ret := _m.Called()

	var r0 entity.CountUsers
	var r1 error
	if rf, ok := ret.Get(0).(func() (entity.CountUsers, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() entity.CountUsers); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entity.CountUsers)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIShorturlRepo_GetAllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUsers'
type MockIShorturlRepo_GetAllUsers_Call struct {
	*mock.Call
}

// GetAllUsers is a helper method to define mock.On call
func (_e *MockIShorturlRepo_Expecter) GetAllUsers() *MockIShorturlRepo_GetAllUsers_Call {
	return &MockIShorturlRepo_GetAllUsers_Call{Call: _e.mock.On("GetAllUsers")}
}

func (_c *MockIShorturlRepo_GetAllUsers_Call) Run(run func()) *MockIShorturlRepo_GetAllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturlRepo_GetAllUsers_Call) Return(_a0 entity.CountUsers, _a1 error) *MockIShorturlRepo_GetAllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIShorturlRepo_GetAllUsers_Call) RunAndReturn(run func() (entity.CountUsers, error)) *MockIShorturlRepo_GetAllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// Post provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturlRepo) Post(_a0 context.Context, _a1 *entity.Shorturl) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturlRepo_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type MockIShorturlRepo_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.Shorturl
func (_e *MockIShorturlRepo_Expecter) Post(_a0 interface{}, _a1 interface{}) *MockIShorturlRepo_Post_Call {
	return &MockIShorturlRepo_Post_Call{Call: _e.mock.On("Post", _a0, _a1)}
}

func (_c *MockIShorturlRepo_Post_Call) Run(run func(_a0 context.Context, _a1 *entity.Shorturl)) *MockIShorturlRepo_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Shorturl))
	})
	return _c
}

func (_c *MockIShorturlRepo_Post_Call) Return(_a0 error) *MockIShorturlRepo_Post_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturlRepo_Post_Call) RunAndReturn(run func(context.Context, *entity.Shorturl) error) *MockIShorturlRepo_Post_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: _a0, _a1
func (_m *MockIShorturlRepo) Put(_a0 context.Context, _a1 *entity.Shorturl) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Shorturl) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturlRepo_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockIShorturlRepo_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *entity.Shorturl
func (_e *MockIShorturlRepo_Expecter) Put(_a0 interface{}, _a1 interface{}) *MockIShorturlRepo_Put_Call {
	return &MockIShorturlRepo_Put_Call{Call: _e.mock.On("Put", _a0, _a1)}
}

func (_c *MockIShorturlRepo_Put_Call) Run(run func(_a0 context.Context, _a1 *entity.Shorturl)) *MockIShorturlRepo_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Shorturl))
	})
	return _c
}

func (_c *MockIShorturlRepo_Put_Call) Return(_a0 error) *MockIShorturlRepo_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturlRepo_Put_Call) RunAndReturn(run func(context.Context, *entity.Shorturl) error) *MockIShorturlRepo_Put_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields:
func (_m *MockIShorturlRepo) Read() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturlRepo_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockIShorturlRepo_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
func (_e *MockIShorturlRepo_Expecter) Read() *MockIShorturlRepo_Read_Call {
	return &MockIShorturlRepo_Read_Call{Call: _e.mock.On("Read")}
}

func (_c *MockIShorturlRepo_Read_Call) Run(run func()) *MockIShorturlRepo_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturlRepo_Read_Call) Return(_a0 error) *MockIShorturlRepo_Read_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturlRepo_Read_Call) RunAndReturn(run func() error) *MockIShorturlRepo_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields:
func (_m *MockIShorturlRepo) Save() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIShorturlRepo_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockIShorturlRepo_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
func (_e *MockIShorturlRepo_Expecter) Save() *MockIShorturlRepo_Save_Call {
	return &MockIShorturlRepo_Save_Call{Call: _e.mock.On("Save")}
}

func (_c *MockIShorturlRepo_Save_Call) Run(run func()) *MockIShorturlRepo_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIShorturlRepo_Save_Call) Return(_a0 error) *MockIShorturlRepo_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIShorturlRepo_Save_Call) RunAndReturn(run func() error) *MockIShorturlRepo_Save_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockIShorturlRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIShorturlRepo creates a new instance of MockIShorturlRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIShorturlRepo(t mockConstructorTestingTNewMockIShorturlRepo) *MockIShorturlRepo {
	mock := &MockIShorturlRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}