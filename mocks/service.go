// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/service.go
//
// Generated by this command:
//
//	mockgen -source=./service/service.go -destination=./mocks/service.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	types "github.com/born2ngopi/example-dolpin/types"
	gomock "go.uber.org/mock/gomock"
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

// Sum mocks base method.
func (m *MockService) Sum(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Sum indicates an expected call of Sum.
func (mr *MockServiceMockRecorder) Sum(a, b any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockService)(nil).Sum), a, b)
}

// SumFromStr mocks base method.
func (m *MockService) SumFromStr(data types.SumField) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumFromStr", data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumFromStr indicates an expected call of SumFromStr.
func (mr *MockServiceMockRecorder) SumFromStr(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumFromStr", reflect.TypeOf((*MockService)(nil).SumFromStr), data)
}
