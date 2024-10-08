// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"
	models "weatherapi/v2/external/models"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryI is a mock of RepositoryI interface.
type MockRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryIMockRecorder
}

// MockRepositoryIMockRecorder is the mock recorder for MockRepositoryI.
type MockRepositoryIMockRecorder struct {
	mock *MockRepositoryI
}

// NewMockRepositoryI creates a new mock instance.
func NewMockRepositoryI(ctrl *gomock.Controller) *MockRepositoryI {
	mock := &MockRepositoryI{ctrl: ctrl}
	mock.recorder = &MockRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryI) EXPECT() *MockRepositoryIMockRecorder {
	return m.recorder
}

// GetGridInfo mocks base method.
func (m *MockRepositoryI) GetGridInfo(ctx context.Context, request models.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGridInfo", ctx, request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGridInfo indicates an expected call of GetGridInfo.
func (mr *MockRepositoryIMockRecorder) GetGridInfo(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGridInfo", reflect.TypeOf((*MockRepositoryI)(nil).GetGridInfo), ctx, request)
}

// SearchWeatherApi mocks base method.
func (m *MockRepositoryI) SearchWeatherApi(ctx context.Context, request models.Request) (WeatherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWeatherApi", ctx, request)
	ret0, _ := ret[0].(WeatherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWeatherApi indicates an expected call of SearchWeatherApi.
func (mr *MockRepositoryIMockRecorder) SearchWeatherApi(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWeatherApi", reflect.TypeOf((*MockRepositoryI)(nil).SearchWeatherApi), ctx, request)
}
