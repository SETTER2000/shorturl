// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/SETTER2000/shorturl/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockShorturl is a mock of Shorturl interface.
type MockShorturl struct {
	ctrl     *gomock.Controller
	recorder *MockShorturlMockRecorder
}

// MockShorturlMockRecorder is the mock recorder for MockShorturl.
type MockShorturlMockRecorder struct {
	mock *MockShorturl
}

// NewMockShorturl creates a new mock instance.
func NewMockShorturl(ctrl *gomock.Controller) *MockShorturl {
	mock := &MockShorturl{ctrl: ctrl}
	mock.recorder = &MockShorturlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShorturl) EXPECT() *MockShorturlMockRecorder {
	return m.recorder
}

// LongLink mocks base method.
func (m *MockShorturl) LongLink(arg0 context.Context, arg1 *entity.Shorturl) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LongLink", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LongLink indicates an expected call of LongLink.
func (mr *MockShorturlMockRecorder) LongLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LongLink", reflect.TypeOf((*MockShorturl)(nil).LongLink), arg0, arg1)
}

// ShortLink mocks base method.
func (m *MockShorturl) ShortLink(arg0 context.Context, arg1 *entity.Shorturl) (*entity.Shorturl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShortLink", arg0, arg1)
	ret0, _ := ret[0].(*entity.Shorturl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShortLink indicates an expected call of ShortLink.
func (mr *MockShorturlMockRecorder) ShortLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShortLink", reflect.TypeOf((*MockShorturl)(nil).ShortLink), arg0, arg1)
}

// Shorten mocks base method.
func (m *MockShorturl) Shorten(arg0 context.Context, arg1 *entity.Shorturl) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shorten", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shorten indicates an expected call of Shorten.
func (mr *MockShorturlMockRecorder) Shorten(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shorten", reflect.TypeOf((*MockShorturl)(nil).Shorten), arg0, arg1)
}

// UserAllLink mocks base method.
func (m *MockShorturl) UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAllLink", ctx, u)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserAllLink indicates an expected call of UserAllLink.
func (mr *MockShorturlMockRecorder) UserAllLink(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAllLink", reflect.TypeOf((*MockShorturl)(nil).UserAllLink), ctx, u)
}

// MockShorturlRepo is a mock of ShorturlRepo interface.
type MockShorturlRepo struct {
	ctrl     *gomock.Controller
	recorder *MockShorturlRepoMockRecorder
}

// MockShorturlRepoMockRecorder is the mock recorder for MockShorturlRepo.
type MockShorturlRepoMockRecorder struct {
	mock *MockShorturlRepo
}

// NewMockShorturlRepo creates a new mock instance.
func NewMockShorturlRepo(ctrl *gomock.Controller) *MockShorturlRepo {
	mock := &MockShorturlRepo{ctrl: ctrl}
	mock.recorder = &MockShorturlRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShorturlRepo) EXPECT() *MockShorturlRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockShorturlRepo) Get(arg0 context.Context, arg1 *entity.Shorturl) (*entity.Shorturl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.Shorturl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShorturlRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShorturlRepo)(nil).Get), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockShorturlRepo) GetAll(arg0 context.Context, arg1 *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockShorturlRepoMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockShorturlRepo)(nil).GetAll), arg0, arg1)
}

// Post mocks base method.
func (m *MockShorturlRepo) Post(arg0 context.Context, arg1 *entity.Shorturl) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockShorturlRepoMockRecorder) Post(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockShorturlRepo)(nil).Post), arg0, arg1)
}

// Put mocks base method.
func (m *MockShorturlRepo) Put(arg0 context.Context, arg1 *entity.Shorturl) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockShorturlRepoMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockShorturlRepo)(nil).Put), arg0, arg1)
}
