// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/client/schedules/openapi.go

// Package schedules is a generated GoMock package.
package schedules

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpRequestDoer is a mock of HttpRequestDoer interface.
type MockHttpRequestDoer struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRequestDoerMockRecorder
}

// MockHttpRequestDoerMockRecorder is the mock recorder for MockHttpRequestDoer.
type MockHttpRequestDoerMockRecorder struct {
	mock *MockHttpRequestDoer
}

// NewMockHttpRequestDoer creates a new mock instance.
func NewMockHttpRequestDoer(ctrl *gomock.Controller) *MockHttpRequestDoer {
	mock := &MockHttpRequestDoer{ctrl: ctrl}
	mock.recorder = &MockHttpRequestDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpRequestDoer) EXPECT() *MockHttpRequestDoerMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHttpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHttpRequestDoerMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpRequestDoer)(nil).Do), req)
}

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// DeleteSchedulesId mocks base method.
func (m *MockClientInterface) DeleteSchedulesId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSchedulesId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSchedulesId indicates an expected call of DeleteSchedulesId.
func (mr *MockClientInterfaceMockRecorder) DeleteSchedulesId(ctx, id interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSchedulesId", reflect.TypeOf((*MockClientInterface)(nil).DeleteSchedulesId), varargs...)
}

// GetSchedulesId mocks base method.
func (m *MockClientInterface) GetSchedulesId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchedulesId", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedulesId indicates an expected call of GetSchedulesId.
func (mr *MockClientInterfaceMockRecorder) GetSchedulesId(ctx, id interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulesId", reflect.TypeOf((*MockClientInterface)(nil).GetSchedulesId), varargs...)
}

// PostSchedules mocks base method.
func (m *MockClientInterface) PostSchedules(ctx context.Context, body PostSchedulesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostSchedules", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSchedules indicates an expected call of PostSchedules.
func (mr *MockClientInterfaceMockRecorder) PostSchedules(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSchedules", reflect.TypeOf((*MockClientInterface)(nil).PostSchedules), varargs...)
}

// PostSchedulesWithBody mocks base method.
func (m *MockClientInterface) PostSchedulesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostSchedulesWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSchedulesWithBody indicates an expected call of PostSchedulesWithBody.
func (mr *MockClientInterfaceMockRecorder) PostSchedulesWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSchedulesWithBody", reflect.TypeOf((*MockClientInterface)(nil).PostSchedulesWithBody), varargs...)
}

// MockClientWithResponsesInterface is a mock of ClientWithResponsesInterface interface.
type MockClientWithResponsesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientWithResponsesInterfaceMockRecorder
}

// MockClientWithResponsesInterfaceMockRecorder is the mock recorder for MockClientWithResponsesInterface.
type MockClientWithResponsesInterfaceMockRecorder struct {
	mock *MockClientWithResponsesInterface
}

// NewMockClientWithResponsesInterface creates a new mock instance.
func NewMockClientWithResponsesInterface(ctrl *gomock.Controller) *MockClientWithResponsesInterface {
	mock := &MockClientWithResponsesInterface{ctrl: ctrl}
	mock.recorder = &MockClientWithResponsesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientWithResponsesInterface) EXPECT() *MockClientWithResponsesInterfaceMockRecorder {
	return m.recorder
}

// DeleteSchedulesIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) DeleteSchedulesIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteSchedulesIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSchedulesIdWithResponse", varargs...)
	ret0, _ := ret[0].(*DeleteSchedulesIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSchedulesIdWithResponse indicates an expected call of DeleteSchedulesIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) DeleteSchedulesIdWithResponse(ctx, id interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSchedulesIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).DeleteSchedulesIdWithResponse), varargs...)
}

// GetSchedulesIdWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetSchedulesIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetSchedulesIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchedulesIdWithResponse", varargs...)
	ret0, _ := ret[0].(*GetSchedulesIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedulesIdWithResponse indicates an expected call of GetSchedulesIdWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetSchedulesIdWithResponse(ctx, id interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulesIdWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetSchedulesIdWithResponse), varargs...)
}

// PostSchedulesWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostSchedulesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSchedulesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostSchedulesWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*PostSchedulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSchedulesWithBodyWithResponse indicates an expected call of PostSchedulesWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostSchedulesWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSchedulesWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostSchedulesWithBodyWithResponse), varargs...)
}

// PostSchedulesWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) PostSchedulesWithResponse(ctx context.Context, body PostSchedulesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSchedulesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostSchedulesWithResponse", varargs...)
	ret0, _ := ret[0].(*PostSchedulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSchedulesWithResponse indicates an expected call of PostSchedulesWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) PostSchedulesWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSchedulesWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).PostSchedulesWithResponse), varargs...)
}
