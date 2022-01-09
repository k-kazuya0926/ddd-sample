// Code generated by MockGen. DO NOT EDIT.
// Source: task_factory.go

// Package mock_task is a generated GoMock package.
package mock_task

import (
	task "ddd-sample/domain/task"
	user "ddd-sample/domain/user"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskFactory is a mock of TaskFactory interface.
type MockTaskFactory struct {
	ctrl     *gomock.Controller
	recorder *MockTaskFactoryMockRecorder
}

// MockTaskFactoryMockRecorder is the mock recorder for MockTaskFactory.
type MockTaskFactoryMockRecorder struct {
	mock *MockTaskFactory
}

// NewMockTaskFactory creates a new mock instance.
func NewMockTaskFactory(ctrl *gomock.Controller) *MockTaskFactory {
	mock := &MockTaskFactory{ctrl: ctrl}
	mock.recorder = &MockTaskFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskFactory) EXPECT() *MockTaskFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskFactory) Create(taskName task.TaskName, dueDate time.Time, userID user.UserID) task.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", taskName, dueDate, userID)
	ret0, _ := ret[0].(task.Task)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTaskFactoryMockRecorder) Create(taskName, dueDate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskFactory)(nil).Create), taskName, dueDate, userID)
}
