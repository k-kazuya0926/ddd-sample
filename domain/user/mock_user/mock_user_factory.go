// Code generated by MockGen. DO NOT EDIT.
// Source: user_factory.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	user "ddd-sample/domain/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserFactory is a mock of UserFactory interface.
type MockUserFactory struct {
	ctrl     *gomock.Controller
	recorder *MockUserFactoryMockRecorder
}

// MockUserFactoryMockRecorder is the mock recorder for MockUserFactory.
type MockUserFactoryMockRecorder struct {
	mock *MockUserFactory
}

// NewMockUserFactory creates a new mock instance.
func NewMockUserFactory(ctrl *gomock.Controller) *MockUserFactory {
	mock := &MockUserFactory{ctrl: ctrl}
	mock.recorder = &MockUserFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserFactory) EXPECT() *MockUserFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserFactory) Create(name user.UserName) user.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(user.User)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserFactoryMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserFactory)(nil).Create), name)
}