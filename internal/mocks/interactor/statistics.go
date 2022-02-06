// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/interactor/statistics.go

// Package mock_interactor is a generated GoMock package.
package mock_interactor

import (
	reflect "reflect"

	entity "github.com/Goalt/hostmonitor/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockStatisticsI is a mock of StatisticsI interface.
type MockStatisticsI struct {
	ctrl     *gomock.Controller
	recorder *MockStatisticsIMockRecorder
}

// MockStatisticsIMockRecorder is the mock recorder for MockStatisticsI.
type MockStatisticsIMockRecorder struct {
	mock *MockStatisticsI
}

// NewMockStatisticsI creates a new mock instance.
func NewMockStatisticsI(ctrl *gomock.Controller) *MockStatisticsI {
	mock := &MockStatisticsI{ctrl: ctrl}
	mock.recorder = &MockStatisticsIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatisticsI) EXPECT() *MockStatisticsIMockRecorder {
	return m.recorder
}

// GetLast mocks base method.
func (m *MockStatisticsI) GetLast() entity.Statistics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLast")
	ret0, _ := ret[0].(entity.Statistics)
	return ret0
}

// GetLast indicates an expected call of GetLast.
func (mr *MockStatisticsIMockRecorder) GetLast() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLast", reflect.TypeOf((*MockStatisticsI)(nil).GetLast))
}

// Update mocks base method.
func (m *MockStatisticsI) Update(arg0 entity.Statistics) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", arg0)
}

// Update indicates an expected call of Update.
func (mr *MockStatisticsIMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatisticsI)(nil).Update), arg0)
}
