// Code generated by MockGen. DO NOT EDIT.
// Source: user_interactor.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/api/domain"
	reflect "reflect"
)

// MockUserInteractor is a mock of UserInteractor interface
type MockUserInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockUserInteractorMockRecorder
}

// MockUserInteractorMockRecorder is the mock recorder for MockUserInteractor
type MockUserInteractorMockRecorder struct {
	mock *MockUserInteractor
}

// NewMockUserInteractor creates a new mock instance
func NewMockUserInteractor(ctrl *gomock.Controller) *MockUserInteractor {
	mock := &MockUserInteractor{ctrl: ctrl}
	mock.recorder = &MockUserInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserInteractor) EXPECT() *MockUserInteractorMockRecorder {
	return m.recorder
}

// UpdatePassword mocks base method
func (m *MockUserInteractor) UpdatePassword(email, currentPassword, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", email, currentPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockUserInteractorMockRecorder) UpdatePassword(email, currentPassword, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserInteractor)(nil).UpdatePassword), email, currentPassword, newPassword)
}

// UpdateProfile mocks base method
func (m *MockUserInteractor) UpdateProfile(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUserInteractorMockRecorder) UpdateProfile(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserInteractor)(nil).UpdateProfile), user)
}
