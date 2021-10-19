// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/thirdpartyrepo.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/redhatinsights/edge-api/pkg/models"
	reflect "reflect"
)

// MockThirdPartyRepoServiceInterface is a mock of ThirdPartyRepoServiceInterface interface
type MockThirdPartyRepoServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockThirdPartyRepoServiceInterfaceMockRecorder
}

// MockThirdPartyRepoServiceInterfaceMockRecorder is the mock recorder for MockThirdPartyRepoServiceInterface
type MockThirdPartyRepoServiceInterfaceMockRecorder struct {
	mock *MockThirdPartyRepoServiceInterface
}

// NewMockThirdPartyRepoServiceInterface creates a new mock instance
func NewMockThirdPartyRepoServiceInterface(ctrl *gomock.Controller) *MockThirdPartyRepoServiceInterface {
	mock := &MockThirdPartyRepoServiceInterface{ctrl: ctrl}
	mock.recorder = &MockThirdPartyRepoServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThirdPartyRepoServiceInterface) EXPECT() *MockThirdPartyRepoServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateThirdPartyRepo mocks base method
func (m *MockThirdPartyRepoServiceInterface) CreateThirdPartyRepo(tprepo *models.ThirdPartyRepo, account string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateThirdPartyRepo", tprepo, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateThirdPartyRepo indicates an expected call of CreateThirdPartyRepo
func (mr *MockThirdPartyRepoServiceInterfaceMockRecorder) CreateThirdPartyRepo(tprepo, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateThirdPartyRepo", reflect.TypeOf((*MockThirdPartyRepoServiceInterface)(nil).CreateThirdPartyRepo), tprepo, account)
}
