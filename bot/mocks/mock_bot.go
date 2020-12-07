// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/HETIC-MT-P2021/GO_TODO_Groupe07/gotodo (interfaces: GoToDo)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGoToDo is a mock of GoToDo interface
type MockGoToDo struct {
	ctrl     *gomock.Controller
	recorder *MockGoToDoMockRecorder
}

// MockGoToDoMockRecorder is the mock recorder for MockGoToDo
type MockGoToDoMockRecorder struct {
	mock *MockGoToDo
}

// NewMockGoToDo creates a new mock instance
func NewMockGoToDo(ctrl *gomock.Controller) *MockGoToDo {
	mock := &MockGoToDo{ctrl: ctrl}
	mock.recorder = &MockGoToDoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGoToDo) EXPECT() *MockGoToDoMockRecorder {
	return m.recorder
}

// MessageReactionAdd mocks base method
func (m *MockGoToDo) MessageReactionAdd(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageReactionAdd", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MessageReactionAdd indicates an expected call of MessageReactionAdd
func (mr *MockGoToDoMockRecorder) MessageReactionAdd(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageReactionAdd", reflect.TypeOf((*MockGoToDo)(nil).MessageReactionAdd), arg0, arg1, arg2)
}
