// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package services is a generated GoMock package.
package service

import (
	model "assignments/layeredcar/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
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

// Create mocks base method.
func (m *MockService) Create(c model.Car) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), c)
}

// Delete mocks base method.
func (m *MockService) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), id)
}

// GetByBrand mocks base method.
func (m *MockService) GetByBrand(brand, engine string) ([]model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBrand", brand, engine)
	ret0, _ := ret[0].([]model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBrand indicates an expected call of GetByBrand.
func (mr *MockServiceMockRecorder) GetByBrand(brand, engine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBrand", reflect.TypeOf((*MockService)(nil).GetByBrand), brand, engine)
}

// GetById mocks base method.
func (m *MockService) GetById(id uuid.UUID) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockServiceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockService)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockService) Update(id uuid.UUID, c model.Car) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, c)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(id, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), id, c)
}
