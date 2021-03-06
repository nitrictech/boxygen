// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nitrictech/boxygen/pkg/proto/builder/v1 (interfaces: Builder_AddServer,Builder_ConfigServer,Builder_CopyServer,Builder_RunServer)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockBuilder_AddServer is a mock of Builder_AddServer interface
type MockBuilder_AddServer struct {
	ctrl     *gomock.Controller
	recorder *MockBuilder_AddServerMockRecorder
}

// MockBuilder_AddServerMockRecorder is the mock recorder for MockBuilder_AddServer
type MockBuilder_AddServerMockRecorder struct {
	mock *MockBuilder_AddServer
}

// NewMockBuilder_AddServer creates a new mock instance
func NewMockBuilder_AddServer(ctrl *gomock.Controller) *MockBuilder_AddServer {
	mock := &MockBuilder_AddServer{ctrl: ctrl}
	mock.recorder = &MockBuilder_AddServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder_AddServer) EXPECT() *MockBuilder_AddServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBuilder_AddServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBuilder_AddServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBuilder_AddServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBuilder_AddServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBuilder_AddServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBuilder_AddServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBuilder_AddServer) Send(arg0 *v1.OutputResponse) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBuilder_AddServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBuilder_AddServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBuilder_AddServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBuilder_AddServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBuilder_AddServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBuilder_AddServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBuilder_AddServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBuilder_AddServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBuilder_AddServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBuilder_AddServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBuilder_AddServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBuilder_AddServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBuilder_AddServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBuilder_AddServer)(nil).SetTrailer), arg0)
}

// MockBuilder_ConfigServer is a mock of Builder_ConfigServer interface
type MockBuilder_ConfigServer struct {
	ctrl     *gomock.Controller
	recorder *MockBuilder_ConfigServerMockRecorder
}

// MockBuilder_ConfigServerMockRecorder is the mock recorder for MockBuilder_ConfigServer
type MockBuilder_ConfigServerMockRecorder struct {
	mock *MockBuilder_ConfigServer
}

// NewMockBuilder_ConfigServer creates a new mock instance
func NewMockBuilder_ConfigServer(ctrl *gomock.Controller) *MockBuilder_ConfigServer {
	mock := &MockBuilder_ConfigServer{ctrl: ctrl}
	mock.recorder = &MockBuilder_ConfigServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder_ConfigServer) EXPECT() *MockBuilder_ConfigServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBuilder_ConfigServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBuilder_ConfigServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBuilder_ConfigServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBuilder_ConfigServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBuilder_ConfigServer) Send(arg0 *v1.OutputResponse) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBuilder_ConfigServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBuilder_ConfigServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBuilder_ConfigServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBuilder_ConfigServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBuilder_ConfigServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBuilder_ConfigServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBuilder_ConfigServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBuilder_ConfigServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBuilder_ConfigServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBuilder_ConfigServer)(nil).SetTrailer), arg0)
}

// MockBuilder_CopyServer is a mock of Builder_CopyServer interface
type MockBuilder_CopyServer struct {
	ctrl     *gomock.Controller
	recorder *MockBuilder_CopyServerMockRecorder
}

// MockBuilder_CopyServerMockRecorder is the mock recorder for MockBuilder_CopyServer
type MockBuilder_CopyServerMockRecorder struct {
	mock *MockBuilder_CopyServer
}

// NewMockBuilder_CopyServer creates a new mock instance
func NewMockBuilder_CopyServer(ctrl *gomock.Controller) *MockBuilder_CopyServer {
	mock := &MockBuilder_CopyServer{ctrl: ctrl}
	mock.recorder = &MockBuilder_CopyServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder_CopyServer) EXPECT() *MockBuilder_CopyServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBuilder_CopyServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBuilder_CopyServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBuilder_CopyServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBuilder_CopyServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBuilder_CopyServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBuilder_CopyServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBuilder_CopyServer) Send(arg0 *v1.OutputResponse) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBuilder_CopyServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBuilder_CopyServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBuilder_CopyServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBuilder_CopyServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBuilder_CopyServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBuilder_CopyServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBuilder_CopyServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBuilder_CopyServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBuilder_CopyServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBuilder_CopyServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBuilder_CopyServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBuilder_CopyServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBuilder_CopyServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBuilder_CopyServer)(nil).SetTrailer), arg0)
}

// MockBuilder_RunServer is a mock of Builder_RunServer interface
type MockBuilder_RunServer struct {
	ctrl     *gomock.Controller
	recorder *MockBuilder_RunServerMockRecorder
}

// MockBuilder_RunServerMockRecorder is the mock recorder for MockBuilder_RunServer
type MockBuilder_RunServerMockRecorder struct {
	mock *MockBuilder_RunServer
}

// NewMockBuilder_RunServer creates a new mock instance
func NewMockBuilder_RunServer(ctrl *gomock.Controller) *MockBuilder_RunServer {
	mock := &MockBuilder_RunServer{ctrl: ctrl}
	mock.recorder = &MockBuilder_RunServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder_RunServer) EXPECT() *MockBuilder_RunServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBuilder_RunServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBuilder_RunServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBuilder_RunServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBuilder_RunServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBuilder_RunServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBuilder_RunServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBuilder_RunServer) Send(arg0 *v1.OutputResponse) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBuilder_RunServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBuilder_RunServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBuilder_RunServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBuilder_RunServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBuilder_RunServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBuilder_RunServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBuilder_RunServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBuilder_RunServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBuilder_RunServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBuilder_RunServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBuilder_RunServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBuilder_RunServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBuilder_RunServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBuilder_RunServer)(nil).SetTrailer), arg0)
}
