// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/YAWAL/ConfRESTcli/api (interfaces: ConfigServiceClient,ConfigService_GetConfigsByTypeClient)

// Package mock_api is a generated GoMock package.
package main

import (
	reflect "reflect"

	api "github.com/YAWAL/GetMeConfAPI/api"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockConfigServiceClient is a mock of ConfigServiceClient interface
type MockConfigServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockConfigServiceClientMockRecorder
}

// MockConfigServiceClientMockRecorder is the mock recorder for MockConfigServiceClient
type MockConfigServiceClientMockRecorder struct {
	mock *MockConfigServiceClient
}

// NewMockConfigServiceClient creates a new mock instance
func NewMockConfigServiceClient(ctrl *gomock.Controller) *MockConfigServiceClient {
	mock := &MockConfigServiceClient{ctrl: ctrl}
	mock.recorder = &MockConfigServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigServiceClient) EXPECT() *MockConfigServiceClientMockRecorder {
	return m.recorder
}

// CreateConfig mocks base method
func (m *MockConfigServiceClient) CreateConfig(arg0 context.Context, arg1 *api.Config, arg2 ...grpc.CallOption) (*api.Responce, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateConfig", varargs...)
	ret0, _ := ret[0].(*api.Responce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfig indicates an expected call of CreateConfig
func (mr *MockConfigServiceClientMockRecorder) CreateConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfig", reflect.TypeOf((*MockConfigServiceClient)(nil).CreateConfig), varargs...)
}

// DeleteConfig mocks base method
func (m *MockConfigServiceClient) DeleteConfig(arg0 context.Context, arg1 *api.DeleteConfigRequest, arg2 ...grpc.CallOption) (*api.Responce, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteConfig", varargs...)
	ret0, _ := ret[0].(*api.Responce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteConfig indicates an expected call of DeleteConfig
func (mr *MockConfigServiceClientMockRecorder) DeleteConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockConfigServiceClient)(nil).DeleteConfig), varargs...)
}

// GetConfigByName mocks base method
func (m *MockConfigServiceClient) GetConfigByName(arg0 context.Context, arg1 *api.GetConfigByNameRequest, arg2 ...grpc.CallOption) (*api.GetConfigResponce, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigByName", varargs...)
	ret0, _ := ret[0].(*api.GetConfigResponce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigByName indicates an expected call of GetConfigByName
func (mr *MockConfigServiceClientMockRecorder) GetConfigByName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigByName", reflect.TypeOf((*MockConfigServiceClient)(nil).GetConfigByName), varargs...)
}

// GetConfigsByType mocks base method
func (m *MockConfigServiceClient) GetConfigsByType(arg0 context.Context, arg1 *api.GetConfigsByTypeRequest, arg2 ...grpc.CallOption) (api.ConfigService_GetConfigsByTypeClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigsByType", varargs...)
	ret0, _ := ret[0].(api.ConfigService_GetConfigsByTypeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigsByType indicates an expected call of GetConfigsByType
func (mr *MockConfigServiceClientMockRecorder) GetConfigsByType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigsByType", reflect.TypeOf((*MockConfigServiceClient)(nil).GetConfigsByType), varargs...)
}

// UpdateConfig mocks base method
func (m *MockConfigServiceClient) UpdateConfig(arg0 context.Context, arg1 *api.Config, arg2 ...grpc.CallOption) (*api.Responce, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateConfig", varargs...)
	ret0, _ := ret[0].(*api.Responce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockConfigServiceClientMockRecorder) UpdateConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockConfigServiceClient)(nil).UpdateConfig), varargs...)
}

// MockConfigService_GetConfigsByTypeClient is a mock of ConfigService_GetConfigsByTypeClient interface
type MockConfigService_GetConfigsByTypeClient struct {
	ctrl     *gomock.Controller
	recorder *MockConfigService_GetConfigsByTypeClientMockRecorder
}

// MockConfigService_GetConfigsByTypeClientMockRecorder is the mock recorder for MockConfigService_GetConfigsByTypeClient
type MockConfigService_GetConfigsByTypeClientMockRecorder struct {
	mock *MockConfigService_GetConfigsByTypeClient
}

// NewMockConfigService_GetConfigsByTypeClient creates a new mock instance
func NewMockConfigService_GetConfigsByTypeClient(ctrl *gomock.Controller) *MockConfigService_GetConfigsByTypeClient {
	mock := &MockConfigService_GetConfigsByTypeClient{ctrl: ctrl}
	mock.recorder = &MockConfigService_GetConfigsByTypeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigService_GetConfigsByTypeClient) EXPECT() *MockConfigService_GetConfigsByTypeClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).Context))
}

// Header mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).Header))
}

// Recv mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) Recv() (*api.GetConfigResponce, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*api.GetConfigResponce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockConfigService_GetConfigsByTypeClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockConfigService_GetConfigsByTypeClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockConfigService_GetConfigsByTypeClient)(nil).Trailer))
}
