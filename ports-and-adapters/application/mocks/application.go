// Code generated by MockGen. DO NOT EDIT.
// Source: application/product.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	application "github.com/codeedu/go-hexagonal/application"
	gomock "github.com/golang/mock/gomock"
)

// MockProductInterface is a mock of ProductInterface interface.
type MockProductInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductInterfaceMockRecorder
}

// MockProductInterfaceMockRecorder is the mock recorder for MockProductInterface.
type MockProductInterfaceMockRecorder struct {
	mock *MockProductInterface
}

// NewMockProductInterface creates a new mock instance.
func NewMockProductInterface(ctrl *gomock.Controller) *MockProductInterface {
	mock := &MockProductInterface{ctrl: ctrl}
	mock.recorder = &MockProductInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductInterface) EXPECT() *MockProductInterfaceMockRecorder {
	return m.recorder
}

// Disable mocks base method.
func (m *MockProductInterface) Disable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockProductInterfaceMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockProductInterface)(nil).Disable))
}

// Enable mocks base method.
func (m *MockProductInterface) Enable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockProductInterfaceMockRecorder) Enable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockProductInterface)(nil).Enable))
}

// GetID mocks base method.
func (m *MockProductInterface) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockProductInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockProductInterface)(nil).GetID))
}

// GetName mocks base method.
func (m *MockProductInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockProductInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockProductInterface)(nil).GetName))
}

// GetPrice mocks base method.
func (m *MockProductInterface) GetPrice() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockProductInterfaceMockRecorder) GetPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockProductInterface)(nil).GetPrice))
}

// GetStatus mocks base method.
func (m *MockProductInterface) GetStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockProductInterfaceMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockProductInterface)(nil).GetStatus))
}

// IsValid mocks base method.
func (m *MockProductInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockProductInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockProductInterface)(nil).IsValid))
}

// MockProductServiceInterface is a mock of ProductServiceInterface interface.
type MockProductServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceInterfaceMockRecorder
}

// MockProductServiceInterfaceMockRecorder is the mock recorder for MockProductServiceInterface.
type MockProductServiceInterfaceMockRecorder struct {
	mock *MockProductServiceInterface
}

// NewMockProductServiceInterface creates a new mock instance.
func NewMockProductServiceInterface(ctrl *gomock.Controller) *MockProductServiceInterface {
	mock := &MockProductServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProductServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceInterface) EXPECT() *MockProductServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductServiceInterface) Create(name string, price float64) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceInterfaceMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductServiceInterface)(nil).Create), name, price)
}

// Disable mocks base method.
func (m *MockProductServiceInterface) Disable(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disable indicates an expected call of Disable.
func (mr *MockProductServiceInterfaceMockRecorder) Disable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockProductServiceInterface)(nil).Disable), product)
}

// Enable mocks base method.
func (m *MockProductServiceInterface) Enable(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enable indicates an expected call of Enable.
func (mr *MockProductServiceInterfaceMockRecorder) Enable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockProductServiceInterface)(nil).Enable), product)
}

// Read mocks base method.
func (m *MockProductServiceInterface) Read(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockProductServiceInterfaceMockRecorder) Read(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProductServiceInterface)(nil).Read), id)
}

// MockProductReader is a mock of ProductReader interface.
type MockProductReader struct {
	ctrl     *gomock.Controller
	recorder *MockProductReaderMockRecorder
}

// MockProductReaderMockRecorder is the mock recorder for MockProductReader.
type MockProductReaderMockRecorder struct {
	mock *MockProductReader
}

// NewMockProductReader creates a new mock instance.
func NewMockProductReader(ctrl *gomock.Controller) *MockProductReader {
	mock := &MockProductReader{ctrl: ctrl}
	mock.recorder = &MockProductReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductReader) EXPECT() *MockProductReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockProductReader) Read(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockProductReaderMockRecorder) Read(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProductReader)(nil).Read), id)
}

// MockProductWriter is a mock of ProductWriter interface.
type MockProductWriter struct {
	ctrl     *gomock.Controller
	recorder *MockProductWriterMockRecorder
}

// MockProductWriterMockRecorder is the mock recorder for MockProductWriter.
type MockProductWriterMockRecorder struct {
	mock *MockProductWriter
}

// NewMockProductWriter creates a new mock instance.
func NewMockProductWriter(ctrl *gomock.Controller) *MockProductWriter {
	mock := &MockProductWriter{ctrl: ctrl}
	mock.recorder = &MockProductWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductWriter) EXPECT() *MockProductWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockProductWriter) Write(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockProductWriterMockRecorder) Write(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockProductWriter)(nil).Write), product)
}

// MockProductPersistenceInterface is a mock of ProductPersistenceInterface interface.
type MockProductPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductPersistenceInterfaceMockRecorder
}

// MockProductPersistenceInterfaceMockRecorder is the mock recorder for MockProductPersistenceInterface.
type MockProductPersistenceInterfaceMockRecorder struct {
	mock *MockProductPersistenceInterface
}

// NewMockProductPersistenceInterface creates a new mock instance.
func NewMockProductPersistenceInterface(ctrl *gomock.Controller) *MockProductPersistenceInterface {
	mock := &MockProductPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockProductPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductPersistenceInterface) EXPECT() *MockProductPersistenceInterfaceMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockProductPersistenceInterface) Read(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockProductPersistenceInterfaceMockRecorder) Read(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProductPersistenceInterface)(nil).Read), id)
}

// Write mocks base method.
func (m *MockProductPersistenceInterface) Write(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockProductPersistenceInterfaceMockRecorder) Write(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockProductPersistenceInterface)(nil).Write), product)
}
