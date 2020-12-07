// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/HETIC-MT-P2021/GO_TODO_Groupe07/reminds (interfaces: Reminds)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	discordgo "github.com/bwmarrin/discordgo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReminds is a mock of Reminds interface
type MockReminds struct {
	ctrl     *gomock.Controller
	recorder *MockRemindsMockRecorder
}

// MockRemindsMockRecorder is the mock recorder for MockReminds
type MockRemindsMockRecorder struct {
	mock *MockReminds
}

// NewMockReminds creates a new mock instance
func NewMockReminds(ctrl *gomock.Controller) *MockReminds {
	mock := &MockReminds{ctrl: ctrl}
	mock.recorder = &MockRemindsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReminds) EXPECT() *MockRemindsMockRecorder {
	return m.recorder
}

// ChannelMessageSend mocks base method
func (m *MockReminds) ChannelMessageSend(arg0, arg1 string) (*discordgo.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelMessageSend", arg0, arg1)
	ret0, _ := ret[0].(*discordgo.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChannelMessageSend indicates an expected call of ChannelMessageSend
func (mr *MockRemindsMockRecorder) ChannelMessageSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelMessageSend", reflect.TypeOf((*MockReminds)(nil).ChannelMessageSend), arg0, arg1)
}

// HandleDefaultCommand mocks base method
func (m *MockReminds) HandleDefaultCommand(arg0 *discordgo.Session, arg1 *discordgo.MessageCreate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleDefaultCommand", arg0, arg1)
}

// HandleDefaultCommand indicates an expected call of HandleDefaultCommand
func (mr *MockRemindsMockRecorder) HandleDefaultCommand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDefaultCommand", reflect.TypeOf((*MockReminds)(nil).HandleDefaultCommand), arg0, arg1)
}

// HandleDeleteRemindCommand mocks base method
func (m *MockReminds) HandleDeleteRemindCommand(arg0 *discordgo.Session, arg1 *discordgo.MessageCreate, arg2 []string, arg3 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleDeleteRemindCommand", arg0, arg1, arg2, arg3)
}

// HandleDeleteRemindCommand indicates an expected call of HandleDeleteRemindCommand
func (mr *MockRemindsMockRecorder) HandleDeleteRemindCommand(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeleteRemindCommand", reflect.TypeOf((*MockReminds)(nil).HandleDeleteRemindCommand), arg0, arg1, arg2, arg3)
}

// HandleGetLastRemindCommand mocks base method
func (m *MockReminds) HandleGetLastRemindCommand(arg0 *discordgo.Session, arg1 *discordgo.MessageCreate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleGetLastRemindCommand", arg0, arg1)
}

// HandleGetLastRemindCommand indicates an expected call of HandleGetLastRemindCommand
func (mr *MockRemindsMockRecorder) HandleGetLastRemindCommand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleGetLastRemindCommand", reflect.TypeOf((*MockReminds)(nil).HandleGetLastRemindCommand), arg0, arg1)
}

// HandleGetRemindsCommand mocks base method
func (m *MockReminds) HandleGetRemindsCommand(arg0 *discordgo.Session, arg1 *discordgo.MessageCreate, arg2 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleGetRemindsCommand", arg0, arg1, arg2)
}

// HandleGetRemindsCommand indicates an expected call of HandleGetRemindsCommand
func (mr *MockRemindsMockRecorder) HandleGetRemindsCommand(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleGetRemindsCommand", reflect.TypeOf((*MockReminds)(nil).HandleGetRemindsCommand), arg0, arg1, arg2)
}

// HandleHelpCommand mocks base method
func (m *MockReminds) HandleHelpCommand(arg0 *discordgo.Session, arg1 *discordgo.MessageCreate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleHelpCommand", arg0, arg1)
}

// HandleHelpCommand indicates an expected call of HandleHelpCommand
func (mr *MockRemindsMockRecorder) HandleHelpCommand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleHelpCommand", reflect.TypeOf((*MockReminds)(nil).HandleHelpCommand), arg0, arg1)
}