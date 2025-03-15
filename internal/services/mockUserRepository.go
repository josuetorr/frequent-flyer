// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/josuetorr/frequent-flyer/internal/services/ (interfaces: UserRepository)
//
// Generated by this command:
//
//	mockgen -destination=internal/services/mockUserRepository.go -package=services github.com/josuetorr/frequent-flyer/internal/services/ UserRepository
//

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	models "github.com/josuetorr/frequent-flyer/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(ctx context.Context, userID string, hard bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, userID, hard)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(ctx, userID, hard any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), ctx, userID, hard)
}

// GetByEmail mocks base method.
func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserRepositoryMockRecorder) GetByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetByEmail), ctx, email)
}

// GetById mocks base method.
func (m *MockUserRepository) GetById(ctx context.Context, userID string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserRepositoryMockRecorder) GetById(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserRepository)(nil).GetById), ctx, userID)
}

// Insert mocks base method.
func (m *MockUserRepository) Insert(ctx context.Context, userID *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockUserRepositoryMockRecorder) Insert(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserRepository)(nil).Insert), ctx, userID)
}

// Update mocks base method.
func (m *MockUserRepository) Update(ctx context.Context, userID string, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(ctx, userID, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), ctx, userID, user)
}
