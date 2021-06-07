// Code generated by MockGen. DO NOT EDIT.
// Source: jwt_generator.go

// Package usecases is a generated GoMock package.
package usecases

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockJWTGenerator is a mock of JWTGenerator interface.
type MockJWTGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockJWTGeneratorMockRecorder
}

// MockJWTGeneratorMockRecorder is the mock recorder for MockJWTGenerator.
type MockJWTGeneratorMockRecorder struct {
	mock *MockJWTGenerator
}

// NewMockJWTGenerator creates a new mock instance.
func NewMockJWTGenerator(ctrl *gomock.Controller) *MockJWTGenerator {
	mock := &MockJWTGenerator{ctrl: ctrl}
	mock.recorder = &MockJWTGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTGenerator) EXPECT() *MockJWTGeneratorMockRecorder {
	return m.recorder
}

// NewToken mocks base method.
func (m *MockJWTGenerator) NewToken(lifetime time.Duration, claims SessionClaims) (string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewToken", lifetime, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewToken indicates an expected call of NewToken.
func (mr *MockJWTGeneratorMockRecorder) NewToken(lifetime, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewToken", reflect.TypeOf((*MockJWTGenerator)(nil).NewToken), lifetime, claims)
}
