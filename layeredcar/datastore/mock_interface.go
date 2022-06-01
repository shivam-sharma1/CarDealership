package datastore

import (
	"assignments/layeredcar/model"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
)

// Mockcar is a mock of car interface.
type Mockcar struct {
	ctrl     *gomock.Controller
	recorder *MockcarMockRecorder
}

// MockcarMockRecorder is a mock recorder for Mockcar.
type MockcarMockRecorder struct {
	mock *Mockcar
}

// NewMockcar will create a new mock instance.
func NewMockcar(ctrl *gomock.Controller) *Mockcar {
	mock := &Mockcar{ctrl: ctrl}
	mock.recorder = &MockcarMockRecorder{mock}
	return mock
}

// EXPECT allows the caller to indicate expected use.
func (m *Mockcar) EXPECT() *MockcarMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *Mockcar) Create(car2 model.Car) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", car2)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockcarMockRecorder) Create(car2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockcar)(nil).Create), car2)
}

// Delete mocks base method.
func (m *Mockcar) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockcarMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mockcar)(nil).Delete), id)
}

// GetByBrand mocks base method getbybrand
func (m *Mockcar) GetByBrand(brand string) ([]model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBrand", brand)
	ret0, _ := ret[0].([]model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBrand indicates an expected call of GetByBrand.
func (mr *MockcarMockRecorder) GetByBrand(brand interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBrand", reflect.TypeOf((*Mockcar)(nil).GetByBrand), brand)
}

// GetById mocks base method getbyid
func (m *Mockcar) GetById(id uuid.UUID) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockcarMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*Mockcar)(nil).GetById), id)
}

// Update mocks base method put or update
func (m *Mockcar) Update(id uuid.UUID, car2 model.Car) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, car2)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockcarMockRecorder) Update(id, car2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Mockcar)(nil).Update), id, car2)
}

// Mockengine is a mock of engine interface.
type Mockengine struct {
	ctrl     *gomock.Controller
	recorder *MockengineMockRecorder
}

// MockengineMockRecorder is the mock recorder for Mockengine.
type MockengineMockRecorder struct {
	mock *Mockengine
}

// NewMockengine creates a new mock instance.
func NewMockengine(ctrl *gomock.Controller) *Mockengine {
	mock := &Mockengine{ctrl: ctrl}
	mock.recorder = &MockengineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockengine) EXPECT() *MockengineMockRecorder {
	return m.recorder
}

// Create mocks base method post or create
func (m *Mockengine) Create(car model.Engine) (model.Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", car)
	ret0, _ := ret[0].(model.Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockengineMockRecorder) Create(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockengine)(nil).Create), car)
}

// Delete mocks base method delete
func (m *Mockengine) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockengineMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mockengine)(nil).Delete), id)
}

// GetById mocks base method getbyid
func (m *Mockengine) GetById(id uuid.UUID) (model.Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockengineMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*Mockengine)(nil).GetById), id)
}

// Update mocks base method put or update
func (m *Mockengine) Update(id uuid.UUID, engine model.Engine) (model.Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, engine)
	ret0, _ := ret[0].(model.Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockengineMockRecorder) Update(id, engine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Mockengine)(nil).Update), id, engine)
}
