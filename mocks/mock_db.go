// Code generated by MockGen. DO NOT EDIT.
// Source: db.go

// Package mock_fmk is a generated GoMock package.
package mock_fmk

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fmk "github.com/opposite-bracket/fmk"
)

// MockIDb is a mock of IDb interface.
type MockIDb struct {
	ctrl     *gomock.Controller
	recorder *MockIDbMockRecorder
}

// MockIDbMockRecorder is the mock recorder for MockIDb.
type MockIDbMockRecorder struct {
	mock *MockIDb
}

// NewMockIDb creates a new mock instance.
func NewMockIDb(ctrl *gomock.Controller) *MockIDb {
	mock := &MockIDb{ctrl: ctrl}
	mock.recorder = &MockIDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDb) EXPECT() *MockIDbMockRecorder {
	return m.recorder
}

// Disconnect mocks base method.
func (m *MockIDb) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockIDbMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockIDb)(nil).Disconnect))
}

// GetModel mocks base method.
func (m *MockIDb) GetModel(colName string) fmk.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModel", colName)
	ret0, _ := ret[0].(fmk.IModel)
	return ret0
}

// GetModel indicates an expected call of GetModel.
func (mr *MockIDbMockRecorder) GetModel(colName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModel", reflect.TypeOf((*MockIDb)(nil).GetModel), colName)
}

// MockIModel is a mock of IModel interface.
type MockIModel struct {
	ctrl     *gomock.Controller
	recorder *MockIModelMockRecorder
}

// MockIModelMockRecorder is the mock recorder for MockIModel.
type MockIModelMockRecorder struct {
	mock *MockIModel
}

// NewMockIModel creates a new mock instance.
func NewMockIModel(ctrl *gomock.Controller) *MockIModel {
	mock := &MockIModel{ctrl: ctrl}
	mock.recorder = &MockIModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIModel) EXPECT() *MockIModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIModel) Delete(filter interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIModelMockRecorder) Delete(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIModel)(nil).Delete), filter)
}

// FindByFilter mocks base method.
func (m *MockIModel) FindByFilter(filter, doc interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByFilter", filter, doc)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByFilter indicates an expected call of FindByFilter.
func (mr *MockIModelMockRecorder) FindByFilter(filter, doc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByFilter", reflect.TypeOf((*MockIModel)(nil).FindByFilter), filter, doc)
}

// Insert mocks base method.
func (m *MockIModel) Insert(doc interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", doc)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockIModelMockRecorder) Insert(doc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIModel)(nil).Insert), doc)
}

// Update mocks base method.
func (m *MockIModel) Update(filter, toChange interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", filter, toChange)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIModelMockRecorder) Update(filter, toChange interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIModel)(nil).Update), filter, toChange)
}
