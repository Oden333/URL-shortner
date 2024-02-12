// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockURLService is a mock of URLService interface.
type MockURLService struct {
	ctrl     *gomock.Controller
	recorder *MockURLServiceMockRecorder
}

// MockURLServiceMockRecorder is the mock recorder for MockURLService.
type MockURLServiceMockRecorder struct {
	mock *MockURLService
}

// NewMockURLService creates a new mock instance.
func NewMockURLService(ctrl *gomock.Controller) *MockURLService {
	mock := &MockURLService{ctrl: ctrl}
	mock.recorder = &MockURLServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLService) EXPECT() *MockURLServiceMockRecorder {
	return m.recorder
}

// GetByAlias mocks base method.
func (m *MockURLService) GetByAlias(alias string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAlias", alias)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAlias indicates an expected call of GetByAlias.
func (mr *MockURLServiceMockRecorder) GetByAlias(alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAlias", reflect.TypeOf((*MockURLService)(nil).GetByAlias), alias)
}

// SaveURL mocks base method.
func (m *MockURLService) SaveURL(url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveURL", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveURL indicates an expected call of SaveURL.
func (mr *MockURLServiceMockRecorder) SaveURL(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveURL", reflect.TypeOf((*MockURLService)(nil).SaveURL), url)
}
