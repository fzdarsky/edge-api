// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/clients/imagebuilder/client.go

// Package mock_imagebuilder is a generated GoMock package.
package mock_imagebuilder

import (
	gomock "github.com/golang/mock/gomock"
	imagebuilder "github.com/redhatinsights/edge-api/pkg/clients/imagebuilder"
	models "github.com/redhatinsights/edge-api/pkg/models"
	http "net/http"
	reflect "reflect"
)

// MockClientInterface is a mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// ComposeCommit mocks base method
func (m *MockClientInterface) ComposeCommit(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComposeCommit", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComposeCommit indicates an expected call of ComposeCommit
func (mr *MockClientInterfaceMockRecorder) ComposeCommit(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComposeCommit", reflect.TypeOf((*MockClientInterface)(nil).ComposeCommit), image)
}

// ComposeInstaller mocks base method
func (m *MockClientInterface) ComposeInstaller(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComposeInstaller", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComposeInstaller indicates an expected call of ComposeInstaller
func (mr *MockClientInterfaceMockRecorder) ComposeInstaller(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComposeInstaller", reflect.TypeOf((*MockClientInterface)(nil).ComposeInstaller), image)
}

// GetCommitStatus mocks base method
func (m *MockClientInterface) GetCommitStatus(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitStatus", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitStatus indicates an expected call of GetCommitStatus
func (mr *MockClientInterfaceMockRecorder) GetCommitStatus(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitStatus", reflect.TypeOf((*MockClientInterface)(nil).GetCommitStatus), image)
}

// GetInstallerStatus mocks base method
func (m *MockClientInterface) GetInstallerStatus(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallerStatus", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallerStatus indicates an expected call of GetInstallerStatus
func (mr *MockClientInterfaceMockRecorder) GetInstallerStatus(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallerStatus", reflect.TypeOf((*MockClientInterface)(nil).GetInstallerStatus), image)
}

// GetMetadata mocks base method
func (m *MockClientInterface) GetMetadata(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata
func (mr *MockClientInterfaceMockRecorder) GetMetadata(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockClientInterface)(nil).GetMetadata), image)
}

// SearchPackage mocks base method
func (m *MockClientInterface) SearchPackage(packageName, arch, dist string) (*models.SearchPackageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPackage", packageName, arch, dist)
	ret0, _ := ret[0].(*models.SearchPackageResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPackage indicates an expected call of SearchPackage
func (mr *MockClientInterfaceMockRecorder) SearchPackage(packageName, arch, dist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPackage", reflect.TypeOf((*MockClientInterface)(nil).SearchPackage), packageName, arch, dist)
}

// ValidatePackages mocks base method
func (m *MockClientInterface) ValidatePackages(pkg imagebuilder.InstalledPackage) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePackages", pkg)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatePackages indicates an expected call of ValidatePackages
func (mr *MockClientInterfaceMockRecorder) ValidatePackages(pkg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePackages", reflect.TypeOf((*MockClientInterface)(nil).ValidatePackages), pkg)
}

// MockHTTPClient is a mock of HTTPClient interface
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockHTTPClientMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), req)
}
