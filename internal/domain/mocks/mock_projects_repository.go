// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mrsufgi/projects-manager/internal/domain (interfaces: ProjectsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/mrsufgi/projects-manager/internal/domain"
	reflect "reflect"
)

// MockProjectsRepository is a mock of ProjectsRepository interface
type MockProjectsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsRepositoryMockRecorder
}

// MockProjectsRepositoryMockRecorder is the mock recorder for MockProjectsRepository
type MockProjectsRepositoryMockRecorder struct {
	mock *MockProjectsRepository
}

// NewMockProjectsRepository creates a new mock instance
func NewMockProjectsRepository(ctrl *gomock.Controller) *MockProjectsRepository {
	mock := &MockProjectsRepository{ctrl: ctrl}
	mock.recorder = &MockProjectsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectsRepository) EXPECT() *MockProjectsRepositoryMockRecorder {
	return m.recorder
}

// CreateProject mocks base method
func (m *MockProjectsRepository) CreateProject(arg0 domain.Project) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockProjectsRepositoryMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectsRepository)(nil).CreateProject), arg0)
}

// DeleteProject mocks base method
func (m *MockProjectsRepository) DeleteProject(arg0 int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockProjectsRepositoryMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectsRepository)(nil).DeleteProject), arg0)
}

// ReadProject mocks base method
func (m *MockProjectsRepository) ReadProject(arg0 int) (*domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProject", arg0)
	ret0, _ := ret[0].(*domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProject indicates an expected call of ReadProject
func (mr *MockProjectsRepositoryMockRecorder) ReadProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProject", reflect.TypeOf((*MockProjectsRepository)(nil).ReadProject), arg0)
}

// SearchProjects mocks base method
func (m *MockProjectsRepository) SearchProjects() (*[]domain.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProjects")
	ret0, _ := ret[0].(*[]domain.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProjects indicates an expected call of SearchProjects
func (mr *MockProjectsRepositoryMockRecorder) SearchProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProjects", reflect.TypeOf((*MockProjectsRepository)(nil).SearchProjects))
}

// UpdateProject mocks base method
func (m *MockProjectsRepository) UpdateProject(arg0 int, arg1 domain.Project) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject
func (mr *MockProjectsRepositoryMockRecorder) UpdateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectsRepository)(nil).UpdateProject), arg0, arg1)
}
