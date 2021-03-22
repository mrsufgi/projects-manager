// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mrsufgi/projects-manager/internal/domain (interfaces: EventsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/mrsufgi/projects-manager/internal/domain"
	reflect "reflect"
)

// MockEventsService is a mock of EventsService interface
type MockEventsService struct {
	ctrl     *gomock.Controller
	recorder *MockEventsServiceMockRecorder
}

// MockEventsServiceMockRecorder is the mock recorder for MockEventsService
type MockEventsServiceMockRecorder struct {
	mock *MockEventsService
}

// NewMockEventsService creates a new mock instance
func NewMockEventsService(ctrl *gomock.Controller) *MockEventsService {
	mock := &MockEventsService{ctrl: ctrl}
	mock.recorder = &MockEventsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventsService) EXPECT() *MockEventsServiceMockRecorder {
	return m.recorder
}

// LogEvent mocks base method
func (m *MockEventsService) LogEvent(arg0 domain.LogEventInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogEvent", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogEvent indicates an expected call of LogEvent
func (mr *MockEventsServiceMockRecorder) LogEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogEvent", reflect.TypeOf((*MockEventsService)(nil).LogEvent), arg0)
}

// ReadEvent mocks base method
func (m *MockEventsService) ReadEvent(arg0 int) (*domain.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadEvent", arg0)
	ret0, _ := ret[0].(*domain.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadEvent indicates an expected call of ReadEvent
func (mr *MockEventsServiceMockRecorder) ReadEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadEvent", reflect.TypeOf((*MockEventsService)(nil).ReadEvent), arg0)
}

// SearchEvents mocks base method
func (m *MockEventsService) SearchEvents(arg0 domain.SearchEventsInput) (*[]domain.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchEvents", arg0)
	ret0, _ := ret[0].(*[]domain.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEvents indicates an expected call of SearchEvents
func (mr *MockEventsServiceMockRecorder) SearchEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEvents", reflect.TypeOf((*MockEventsService)(nil).SearchEvents), arg0)
}
