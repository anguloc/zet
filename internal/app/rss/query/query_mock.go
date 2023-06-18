// Code generated by MockGen. DO NOT EDIT.
// Source: ./query.go

// Package query is a generated GoMock package.
package query

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIQuery is a mock of IQuery interface.
type MockIQuery struct {
	ctrl     *gomock.Controller
	recorder *MockIQueryMockRecorder
}

// MockIQueryMockRecorder is the mock recorder for MockIQuery.
type MockIQueryMockRecorder struct {
	mock *MockIQuery
}

// NewMockIQuery creates a new mock instance.
func NewMockIQuery(ctrl *gomock.Controller) *MockIQuery {
	mock := &MockIQuery{ctrl: ctrl}
	mock.recorder = &MockIQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIQuery) EXPECT() *MockIQueryMockRecorder {
	return m.recorder
}

// Command mocks base method.
func (m *MockIQuery) Command(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Command indicates an expected call of Command.
func (mr *MockIQueryMockRecorder) Command(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockIQuery)(nil).Command), ctx)
}

// Get mocks base method.
func (m *MockIQuery) Get(ctx context.Context) (*Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(*Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIQueryMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIQuery)(nil).Get), ctx)
}
