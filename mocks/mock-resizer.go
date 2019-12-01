// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Dimitriy14/image-resizing/usecases (interfaces: ImageResizer)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/Dimitriy14/image-resizing/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockResizer is a mock of ImageResizer interface
type MockResizer struct {
	ctrl     *gomock.Controller
	recorder *MockResizerMockRecorder
}

// MockResizerMockRecorder is the mock recorder for MockResizer
type MockResizerMockRecorder struct {
	mock *MockResizer
}

// NewMockResizer creates a new mock instance
func NewMockResizer(ctrl *gomock.Controller) *MockResizer {
	mock := &MockResizer{ctrl: ctrl}
	mock.recorder = &MockResizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResizer) EXPECT() *MockResizerMockRecorder {
	return m.recorder
}

// Resize mocks base method
func (m *MockResizer) Resize(arg0 []byte, arg1 models.ResizeParams) ([]byte, error) {
	ret := m.ctrl.Call(m, "Resize", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize
func (mr *MockResizerMockRecorder) Resize(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockResizer)(nil).Resize), arg0, arg1)
}
