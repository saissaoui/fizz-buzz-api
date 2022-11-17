// Code generated by MockGen. DO NOT EDIT.
// Source: ./statistics.go

// Package statistics is a generated GoMock package.
package statistics

import (
	fizzbuzz "fizz-buzz-api/services/fizzbuzz"
	statistics "fizz-buzz-api/services/statistics"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CountRequest mocks base method.
func (m *MockService) CountRequest(request fizzbuzz.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountRequest", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CountRequest indicates an expected call of CountRequest.
func (mr *MockServiceMockRecorder) CountRequest(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountRequest", reflect.TypeOf((*MockService)(nil).CountRequest), request)
}

// GetStatistics mocks base method.
func (m *MockService) GetStatistics() ([]statistics.RequestStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatistics")
	ret0, _ := ret[0].([]statistics.RequestStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatistics indicates an expected call of GetStatistics.
func (mr *MockServiceMockRecorder) GetStatistics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatistics", reflect.TypeOf((*MockService)(nil).GetStatistics))
}