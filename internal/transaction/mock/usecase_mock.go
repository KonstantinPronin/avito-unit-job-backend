// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/transaction/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	model "github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockUsecase) Add(tr *model.Transaction) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", tr)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockUsecaseMockRecorder) Add(tr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUsecase)(nil).Add), tr)
}

// Get mocks base method
func (m *MockUsecase) Get(tid uint64) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", tid)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsecaseMockRecorder) Get(tid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsecase)(nil).Get), tid)
}

// GetOrderByDate mocks base method
func (m *MockUsecase) GetOrderByDate(uid uint64, page uint, desc bool) (model.History, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByDate", uid, page, desc)
	ret0, _ := ret[0].(model.History)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByDate indicates an expected call of GetOrderByDate
func (mr *MockUsecaseMockRecorder) GetOrderByDate(uid, page, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByDate", reflect.TypeOf((*MockUsecase)(nil).GetOrderByDate), uid, page, desc)
}

// GetOrderBySum mocks base method
func (m *MockUsecase) GetOrderBySum(uid uint64, page uint, desc bool) (model.History, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderBySum", uid, page, desc)
	ret0, _ := ret[0].(model.History)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderBySum indicates an expected call of GetOrderBySum
func (mr *MockUsecaseMockRecorder) GetOrderBySum(uid, page, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderBySum", reflect.TypeOf((*MockUsecase)(nil).GetOrderBySum), uid, page, desc)
}