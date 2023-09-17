// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anguloc/zet/internal/dao/zet_query (interfaces: IRssDo)

// Package zet_query is a generated GoMock package.
package zet_query

import (
	context "context"
	reflect "reflect"

	zet_model "github.com/anguloc/zet/internal/dao/zet_model"
	gomock "github.com/golang/mock/gomock"
	gen "gorm.io/gen"
	field "gorm.io/gen/field"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
	schema "gorm.io/gorm/schema"
)

// MockIRssDo is a mock of IRssDo interface.
type MockIRssDo struct {
	ctrl     *gomock.Controller
	recorder *MockIRssDoMockRecorder
}

// MockIRssDoMockRecorder is the mock recorder for MockIRssDo.
type MockIRssDoMockRecorder struct {
	mock *MockIRssDo
}

// NewMockIRssDo creates a new mock instance.
func NewMockIRssDo(ctrl *gomock.Controller) *MockIRssDo {
	mock := &MockIRssDo{ctrl: ctrl}
	mock.recorder = &MockIRssDoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRssDo) EXPECT() *MockIRssDoMockRecorder {
	return m.recorder
}

// As mocks base method.
func (m *MockIRssDo) As(arg0 string) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "As", arg0)
	ret0, _ := ret[0].(gen.Dao)
	return ret0
}

// As indicates an expected call of As.
func (mr *MockIRssDoMockRecorder) As(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "As", reflect.TypeOf((*MockIRssDo)(nil).As), arg0)
}

// Assign mocks base method.
func (m *MockIRssDo) Assign(arg0 ...field.AssignExpr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Assign", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Assign indicates an expected call of Assign.
func (mr *MockIRssDoMockRecorder) Assign(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*MockIRssDo)(nil).Assign), arg0...)
}

// Attrs mocks base method.
func (m *MockIRssDo) Attrs(arg0 ...field.AssignExpr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Attrs", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Attrs indicates an expected call of Attrs.
func (mr *MockIRssDoMockRecorder) Attrs(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockIRssDo)(nil).Attrs), arg0...)
}

// BeCond mocks base method.
func (m *MockIRssDo) BeCond() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeCond")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// BeCond indicates an expected call of BeCond.
func (mr *MockIRssDoMockRecorder) BeCond() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeCond", reflect.TypeOf((*MockIRssDo)(nil).BeCond))
}

// Clauses mocks base method.
func (m *MockIRssDo) Clauses(arg0 ...clause.Expression) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Clauses", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Clauses indicates an expected call of Clauses.
func (mr *MockIRssDoMockRecorder) Clauses(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clauses", reflect.TypeOf((*MockIRssDo)(nil).Clauses), arg0...)
}

// Columns mocks base method.
func (m *MockIRssDo) Columns(arg0 ...field.Expr) gen.Columns {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Columns", varargs...)
	ret0, _ := ret[0].(gen.Columns)
	return ret0
}

// Columns indicates an expected call of Columns.
func (mr *MockIRssDoMockRecorder) Columns(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockIRssDo)(nil).Columns), arg0...)
}

// CondError mocks base method.
func (m *MockIRssDo) CondError() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CondError")
	ret0, _ := ret[0].(error)
	return ret0
}

// CondError indicates an expected call of CondError.
func (mr *MockIRssDoMockRecorder) CondError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CondError", reflect.TypeOf((*MockIRssDo)(nil).CondError))
}

// Count mocks base method.
func (m *MockIRssDo) Count() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIRssDoMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIRssDo)(nil).Count))
}

// Create mocks base method.
func (m *MockIRssDo) Create(arg0 ...*zet_model.Rss) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIRssDoMockRecorder) Create(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRssDo)(nil).Create), arg0...)
}

// CreateInBatches mocks base method.
func (m *MockIRssDo) CreateInBatches(arg0 []*zet_model.Rss, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInBatches indicates an expected call of CreateInBatches.
func (mr *MockIRssDoMockRecorder) CreateInBatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatches", reflect.TypeOf((*MockIRssDo)(nil).CreateInBatches), arg0, arg1)
}

// Debug mocks base method.
func (m *MockIRssDo) Debug() IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug")
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Debug indicates an expected call of Debug.
func (mr *MockIRssDoMockRecorder) Debug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockIRssDo)(nil).Debug))
}

// Delete mocks base method.
func (m *MockIRssDo) Delete(arg0 ...*zet_model.Rss) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIRssDoMockRecorder) Delete(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRssDo)(nil).Delete), arg0...)
}

// Distinct mocks base method.
func (m *MockIRssDo) Distinct(arg0 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Distinct", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Distinct indicates an expected call of Distinct.
func (mr *MockIRssDoMockRecorder) Distinct(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distinct", reflect.TypeOf((*MockIRssDo)(nil).Distinct), arg0...)
}

// Find mocks base method.
func (m *MockIRssDo) Find() ([]*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find")
	ret0, _ := ret[0].([]*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIRssDoMockRecorder) Find() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIRssDo)(nil).Find))
}

// FindByPage mocks base method.
func (m *MockIRssDo) FindByPage(arg0, arg1 int) ([]*zet_model.Rss, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPage", arg0, arg1)
	ret0, _ := ret[0].([]*zet_model.Rss)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindByPage indicates an expected call of FindByPage.
func (mr *MockIRssDoMockRecorder) FindByPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPage", reflect.TypeOf((*MockIRssDo)(nil).FindByPage), arg0, arg1)
}

// FindInBatch mocks base method.
func (m *MockIRssDo) FindInBatch(arg0 int, arg1 func(gen.Dao, int) error) ([]*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInBatch", arg0, arg1)
	ret0, _ := ret[0].([]*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInBatch indicates an expected call of FindInBatch.
func (mr *MockIRssDoMockRecorder) FindInBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInBatch", reflect.TypeOf((*MockIRssDo)(nil).FindInBatch), arg0, arg1)
}

// FindInBatches mocks base method.
func (m *MockIRssDo) FindInBatches(arg0 *[]*zet_model.Rss, arg1 int, arg2 func(gen.Dao, int) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInBatches", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindInBatches indicates an expected call of FindInBatches.
func (mr *MockIRssDoMockRecorder) FindInBatches(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInBatches", reflect.TypeOf((*MockIRssDo)(nil).FindInBatches), arg0, arg1, arg2)
}

// First mocks base method.
func (m *MockIRssDo) First() (*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// First indicates an expected call of First.
func (mr *MockIRssDoMockRecorder) First() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockIRssDo)(nil).First))
}

// FirstOrCreate mocks base method.
func (m *MockIRssDo) FirstOrCreate() (*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstOrCreate")
	ret0, _ := ret[0].(*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstOrCreate indicates an expected call of FirstOrCreate.
func (mr *MockIRssDoMockRecorder) FirstOrCreate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrCreate", reflect.TypeOf((*MockIRssDo)(nil).FirstOrCreate))
}

// FirstOrInit mocks base method.
func (m *MockIRssDo) FirstOrInit() (*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstOrInit")
	ret0, _ := ret[0].(*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstOrInit indicates an expected call of FirstOrInit.
func (mr *MockIRssDoMockRecorder) FirstOrInit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrInit", reflect.TypeOf((*MockIRssDo)(nil).FirstOrInit))
}

// Group mocks base method.
func (m *MockIRssDo) Group(arg0 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Group", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *MockIRssDoMockRecorder) Group(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockIRssDo)(nil).Group), arg0...)
}

// Having mocks base method.
func (m *MockIRssDo) Having(arg0 ...gen.Condition) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Having", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Having indicates an expected call of Having.
func (mr *MockIRssDoMockRecorder) Having(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Having", reflect.TypeOf((*MockIRssDo)(nil).Having), arg0...)
}

// Join mocks base method.
func (m *MockIRssDo) Join(arg0 schema.Tabler, arg1 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Join", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Join indicates an expected call of Join.
func (mr *MockIRssDoMockRecorder) Join(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockIRssDo)(nil).Join), varargs...)
}

// Joins mocks base method.
func (m *MockIRssDo) Joins(arg0 ...field.RelationField) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Joins", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Joins indicates an expected call of Joins.
func (mr *MockIRssDoMockRecorder) Joins(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Joins", reflect.TypeOf((*MockIRssDo)(nil).Joins), arg0...)
}

// Last mocks base method.
func (m *MockIRssDo) Last() (*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Last indicates an expected call of Last.
func (mr *MockIRssDoMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockIRssDo)(nil).Last))
}

// LeftJoin mocks base method.
func (m *MockIRssDo) LeftJoin(arg0 schema.Tabler, arg1 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LeftJoin", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// LeftJoin indicates an expected call of LeftJoin.
func (mr *MockIRssDoMockRecorder) LeftJoin(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeftJoin", reflect.TypeOf((*MockIRssDo)(nil).LeftJoin), varargs...)
}

// Limit mocks base method.
func (m *MockIRssDo) Limit(arg0 int) IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", arg0)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Limit indicates an expected call of Limit.
func (mr *MockIRssDoMockRecorder) Limit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockIRssDo)(nil).Limit), arg0)
}

// Not mocks base method.
func (m *MockIRssDo) Not(arg0 ...gen.Condition) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Not", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Not indicates an expected call of Not.
func (mr *MockIRssDoMockRecorder) Not(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Not", reflect.TypeOf((*MockIRssDo)(nil).Not), arg0...)
}

// Offset mocks base method.
func (m *MockIRssDo) Offset(arg0 int) IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", arg0)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *MockIRssDoMockRecorder) Offset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockIRssDo)(nil).Offset), arg0)
}

// Omit mocks base method.
func (m *MockIRssDo) Omit(arg0 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Omit", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Omit indicates an expected call of Omit.
func (mr *MockIRssDoMockRecorder) Omit(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Omit", reflect.TypeOf((*MockIRssDo)(nil).Omit), arg0...)
}

// Or mocks base method.
func (m *MockIRssDo) Or(arg0 ...gen.Condition) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Or", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Or indicates an expected call of Or.
func (mr *MockIRssDoMockRecorder) Or(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*MockIRssDo)(nil).Or), arg0...)
}

// Order mocks base method.
func (m *MockIRssDo) Order(arg0 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Order", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Order indicates an expected call of Order.
func (mr *MockIRssDoMockRecorder) Order(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockIRssDo)(nil).Order), arg0...)
}

// Pluck mocks base method.
func (m *MockIRssDo) Pluck(arg0 field.Expr, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pluck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pluck indicates an expected call of Pluck.
func (mr *MockIRssDoMockRecorder) Pluck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pluck", reflect.TypeOf((*MockIRssDo)(nil).Pluck), arg0, arg1)
}

// Preload mocks base method.
func (m *MockIRssDo) Preload(arg0 ...field.RelationField) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Preload", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Preload indicates an expected call of Preload.
func (mr *MockIRssDoMockRecorder) Preload(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preload", reflect.TypeOf((*MockIRssDo)(nil).Preload), arg0...)
}

// ReadDB mocks base method.
func (m *MockIRssDo) ReadDB() IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDB")
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// ReadDB indicates an expected call of ReadDB.
func (mr *MockIRssDoMockRecorder) ReadDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDB", reflect.TypeOf((*MockIRssDo)(nil).ReadDB))
}

// ReplaceDB mocks base method.
func (m *MockIRssDo) ReplaceDB(arg0 *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceDB", arg0)
}

// ReplaceDB indicates an expected call of ReplaceDB.
func (mr *MockIRssDoMockRecorder) ReplaceDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceDB", reflect.TypeOf((*MockIRssDo)(nil).ReplaceDB), arg0)
}

// Returning mocks base method.
func (m *MockIRssDo) Returning(arg0 interface{}, arg1 ...string) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Returning", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Returning indicates an expected call of Returning.
func (mr *MockIRssDoMockRecorder) Returning(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Returning", reflect.TypeOf((*MockIRssDo)(nil).Returning), varargs...)
}

// RightJoin mocks base method.
func (m *MockIRssDo) RightJoin(arg0 schema.Tabler, arg1 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RightJoin", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// RightJoin indicates an expected call of RightJoin.
func (mr *MockIRssDoMockRecorder) RightJoin(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RightJoin", reflect.TypeOf((*MockIRssDo)(nil).RightJoin), varargs...)
}

// Save mocks base method.
func (m *MockIRssDo) Save(arg0 ...*zet_model.Rss) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIRssDoMockRecorder) Save(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRssDo)(nil).Save), arg0...)
}

// Scan mocks base method.
func (m *MockIRssDo) Scan(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockIRssDoMockRecorder) Scan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockIRssDo)(nil).Scan), arg0)
}

// ScanByPage mocks base method.
func (m *MockIRssDo) ScanByPage(arg0 interface{}, arg1, arg2 int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanByPage", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanByPage indicates an expected call of ScanByPage.
func (mr *MockIRssDoMockRecorder) ScanByPage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanByPage", reflect.TypeOf((*MockIRssDo)(nil).ScanByPage), arg0, arg1, arg2)
}

// Scopes mocks base method.
func (m *MockIRssDo) Scopes(arg0 ...func(gen.Dao) gen.Dao) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scopes", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Scopes indicates an expected call of Scopes.
func (mr *MockIRssDoMockRecorder) Scopes(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scopes", reflect.TypeOf((*MockIRssDo)(nil).Scopes), arg0...)
}

// Select mocks base method.
func (m *MockIRssDo) Select(arg0 ...field.Expr) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockIRssDoMockRecorder) Select(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIRssDo)(nil).Select), arg0...)
}

// Session mocks base method.
func (m *MockIRssDo) Session(arg0 *gorm.Session) IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", arg0)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Session indicates an expected call of Session.
func (mr *MockIRssDoMockRecorder) Session(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockIRssDo)(nil).Session), arg0)
}

// TableName mocks base method.
func (m *MockIRssDo) TableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName.
func (mr *MockIRssDoMockRecorder) TableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableName", reflect.TypeOf((*MockIRssDo)(nil).TableName))
}

// Take mocks base method.
func (m *MockIRssDo) Take() (*zet_model.Rss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Take")
	ret0, _ := ret[0].(*zet_model.Rss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Take indicates an expected call of Take.
func (mr *MockIRssDoMockRecorder) Take() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockIRssDo)(nil).Take))
}

// UnderlyingDB mocks base method.
func (m *MockIRssDo) UnderlyingDB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnderlyingDB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// UnderlyingDB indicates an expected call of UnderlyingDB.
func (mr *MockIRssDoMockRecorder) UnderlyingDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnderlyingDB", reflect.TypeOf((*MockIRssDo)(nil).UnderlyingDB))
}

// Unscoped mocks base method.
func (m *MockIRssDo) Unscoped() IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unscoped")
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Unscoped indicates an expected call of Unscoped.
func (mr *MockIRssDoMockRecorder) Unscoped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unscoped", reflect.TypeOf((*MockIRssDo)(nil).Unscoped))
}

// Update mocks base method.
func (m *MockIRssDo) Update(arg0 field.Expr, arg1 interface{}) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIRssDoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRssDo)(nil).Update), arg0, arg1)
}

// UpdateColumn mocks base method.
func (m *MockIRssDo) UpdateColumn(arg0 field.Expr, arg1 interface{}) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", arg0, arg1)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *MockIRssDoMockRecorder) UpdateColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockIRssDo)(nil).UpdateColumn), arg0, arg1)
}

// UpdateColumnSimple mocks base method.
func (m *MockIRssDo) UpdateColumnSimple(arg0 ...field.AssignExpr) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateColumnSimple", varargs...)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateColumnSimple indicates an expected call of UpdateColumnSimple.
func (mr *MockIRssDoMockRecorder) UpdateColumnSimple(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumnSimple", reflect.TypeOf((*MockIRssDo)(nil).UpdateColumnSimple), arg0...)
}

// UpdateColumns mocks base method.
func (m *MockIRssDo) UpdateColumns(arg0 interface{}) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumns", arg0)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateColumns indicates an expected call of UpdateColumns.
func (mr *MockIRssDoMockRecorder) UpdateColumns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumns", reflect.TypeOf((*MockIRssDo)(nil).UpdateColumns), arg0)
}

// UpdateFrom mocks base method.
func (m *MockIRssDo) UpdateFrom(arg0 gen.SubQuery) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFrom", arg0)
	ret0, _ := ret[0].(gen.Dao)
	return ret0
}

// UpdateFrom indicates an expected call of UpdateFrom.
func (mr *MockIRssDoMockRecorder) UpdateFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFrom", reflect.TypeOf((*MockIRssDo)(nil).UpdateFrom), arg0)
}

// UpdateSimple mocks base method.
func (m *MockIRssDo) UpdateSimple(arg0 ...field.AssignExpr) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSimple", varargs...)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSimple indicates an expected call of UpdateSimple.
func (mr *MockIRssDoMockRecorder) UpdateSimple(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSimple", reflect.TypeOf((*MockIRssDo)(nil).UpdateSimple), arg0...)
}

// Updates mocks base method.
func (m *MockIRssDo) Updates(arg0 interface{}) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", arg0)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Updates indicates an expected call of Updates.
func (mr *MockIRssDoMockRecorder) Updates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockIRssDo)(nil).Updates), arg0)
}

// Where mocks base method.
func (m *MockIRssDo) Where(arg0 ...gen.Condition) IRssDo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockIRssDoMockRecorder) Where(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockIRssDo)(nil).Where), arg0...)
}

// WithContext mocks base method.
func (m *MockIRssDo) WithContext(arg0 context.Context) IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", arg0)
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockIRssDoMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockIRssDo)(nil).WithContext), arg0)
}

// WithResult mocks base method.
func (m *MockIRssDo) WithResult(arg0 func(gen.Dao)) gen.ResultInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithResult", arg0)
	ret0, _ := ret[0].(gen.ResultInfo)
	return ret0
}

// WithResult indicates an expected call of WithResult.
func (mr *MockIRssDoMockRecorder) WithResult(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithResult", reflect.TypeOf((*MockIRssDo)(nil).WithResult), arg0)
}

// WriteDB mocks base method.
func (m *MockIRssDo) WriteDB() IRssDo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteDB")
	ret0, _ := ret[0].(IRssDo)
	return ret0
}

// WriteDB indicates an expected call of WriteDB.
func (mr *MockIRssDoMockRecorder) WriteDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDB", reflect.TypeOf((*MockIRssDo)(nil).WriteDB))
}

// underlyingDB mocks base method.
func (m *MockIRssDo) underlyingDB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "underlyingDB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// underlyingDB indicates an expected call of underlyingDB.
func (mr *MockIRssDoMockRecorder) underlyingDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "underlyingDB", reflect.TypeOf((*MockIRssDo)(nil).underlyingDB))
}

// underlyingDO mocks base method.
func (m *MockIRssDo) underlyingDO() *gen.DO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "underlyingDO")
	ret0, _ := ret[0].(*gen.DO)
	return ret0
}

// underlyingDO indicates an expected call of underlyingDO.
func (mr *MockIRssDoMockRecorder) underlyingDO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "underlyingDO", reflect.TypeOf((*MockIRssDo)(nil).underlyingDO))
}
