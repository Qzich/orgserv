// Code generated by MockGen. DO NOT EDIT.
// Source: ../repository/repository.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	users "github.com/qzich/orgserv/entity/users"
	uuid "github.com/qzich/orgserv/pkg/uuid"
)

// MockUsersRepository is a mock of UsersRepository interface.
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository.
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance.
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// GetAuthUser mocks base method.
func (m *MockUsersRepository) GetAuthUser(email string) (users.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthUser", email)
	ret0, _ := ret[0].(users.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthUser indicates an expected call of GetAuthUser.
func (mr *MockUsersRepositoryMockRecorder) GetAuthUser(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthUser", reflect.TypeOf((*MockUsersRepository)(nil).GetAuthUser), email)
}

// GetUserByID mocks base method.
func (m *MockUsersRepository) GetUserByID(userID uuid.UUID) (users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUsersRepositoryMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUsersRepository)(nil).GetUserByID), userID)
}

// InsertUser mocks base method.
func (m *MockUsersRepository) InsertUser(data users.User, passHash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", data, passHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockUsersRepositoryMockRecorder) InsertUser(data, passHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUsersRepository)(nil).InsertUser), data, passHash)
}

// SearchUsers mocks base method.
func (m *MockUsersRepository) SearchUsers() ([]users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers")
	ret0, _ := ret[0].([]users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockUsersRepositoryMockRecorder) SearchUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockUsersRepository)(nil).SearchUsers))
}

// UpdateUser mocks base method.
func (m *MockUsersRepository) UpdateUser(userID uuid.UUID, data users.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userID, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUsersRepositoryMockRecorder) UpdateUser(userID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsersRepository)(nil).UpdateUser), userID, data)
}
