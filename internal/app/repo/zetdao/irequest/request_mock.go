// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package irequest is a generated GoMock package.
package irequest

import (
	context "context"
	reflect "reflect"

	model "github.com/anguloc/zet/internal/app/repo/zetdao/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// FirstByMarkStatus mocks base method.
func (m *MockRepo) FirstByMarkStatus(ctx context.Context, mark string, status int32) (*model.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstByMarkStatus", ctx, mark, status)
	ret0, _ := ret[0].(*model.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstByMarkStatus indicates an expected call of FirstByMarkStatus.
func (mr *MockRepoMockRecorder) FirstByMarkStatus(ctx, mark, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstByMarkStatus", reflect.TypeOf((*MockRepo)(nil).FirstByMarkStatus), ctx, mark, status)
}

// Insert mocks base method.
func (m *MockRepo) Insert(ctx context.Context, data *model.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepoMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepo)(nil).Insert), ctx, data)
}

// UpdateStatusByIds mocks base method.
func (m *MockRepo) UpdateStatusByIds(ctx context.Context, ids []int64, newStatus, oldStatus int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByIds", ctx, ids, newStatus, oldStatus)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatusByIds indicates an expected call of UpdateStatusByIds.
func (mr *MockRepoMockRecorder) UpdateStatusByIds(ctx, ids, newStatus, oldStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByIds", reflect.TypeOf((*MockRepo)(nil).UpdateStatusByIds), ctx, ids, newStatus, oldStatus)
}

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// FirstByMarkStatus mocks base method.
func (m *MockReader) FirstByMarkStatus(ctx context.Context, mark string, status int32) (*model.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstByMarkStatus", ctx, mark, status)
	ret0, _ := ret[0].(*model.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstByMarkStatus indicates an expected call of FirstByMarkStatus.
func (mr *MockReaderMockRecorder) FirstByMarkStatus(ctx, mark, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstByMarkStatus", reflect.TypeOf((*MockReader)(nil).FirstByMarkStatus), ctx, mark, status)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockWriter) Insert(ctx context.Context, data *model.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockWriterMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockWriter)(nil).Insert), ctx, data)
}

// UpdateStatusByIds mocks base method.
func (m *MockWriter) UpdateStatusByIds(ctx context.Context, ids []int64, newStatus, oldStatus int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByIds", ctx, ids, newStatus, oldStatus)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatusByIds indicates an expected call of UpdateStatusByIds.
func (mr *MockWriterMockRecorder) UpdateStatusByIds(ctx, ids, newStatus, oldStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByIds", reflect.TypeOf((*MockWriter)(nil).UpdateStatusByIds), ctx, ids, newStatus, oldStatus)
}