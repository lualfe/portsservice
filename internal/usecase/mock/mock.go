// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/lualfe/portsservice/internal/entity"
	portsstream "github.com/lualfe/portsservice/pkg/portsstream"
)

// MockPortsRepo is a mock of PortsRepo interface.
type MockPortsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPortsRepoMockRecorder
}

// MockPortsRepoMockRecorder is the mock recorder for MockPortsRepo.
type MockPortsRepoMockRecorder struct {
	mock *MockPortsRepo
}

// NewMockPortsRepo creates a new mock instance.
func NewMockPortsRepo(ctrl *gomock.Controller) *MockPortsRepo {
	mock := &MockPortsRepo{ctrl: ctrl}
	mock.recorder = &MockPortsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortsRepo) EXPECT() *MockPortsRepoMockRecorder {
	return m.recorder
}

// BulkUpsertPort mocks base method.
func (m *MockPortsRepo) BulkUpsertPort(ctx context.Context, ports []entity.Port) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpsertPort", ctx, ports)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpsertPort indicates an expected call of BulkUpsertPort.
func (mr *MockPortsRepoMockRecorder) BulkUpsertPort(ctx, ports interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpsertPort", reflect.TypeOf((*MockPortsRepo)(nil).BulkUpsertPort), ctx, ports)
}

// UpsertPort mocks base method.
func (m *MockPortsRepo) UpsertPort(ctx context.Context, port entity.Port) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPort", ctx, port)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPort indicates an expected call of UpsertPort.
func (mr *MockPortsRepoMockRecorder) UpsertPort(ctx, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPort", reflect.TypeOf((*MockPortsRepo)(nil).UpsertPort), ctx, port)
}

// MockPortsStreamer is a mock of PortsStreamer interface.
type MockPortsStreamer struct {
	ctrl     *gomock.Controller
	recorder *MockPortsStreamerMockRecorder
}

// MockPortsStreamerMockRecorder is the mock recorder for MockPortsStreamer.
type MockPortsStreamerMockRecorder struct {
	mock *MockPortsStreamer
}

// NewMockPortsStreamer creates a new mock instance.
func NewMockPortsStreamer(ctrl *gomock.Controller) *MockPortsStreamer {
	mock := &MockPortsStreamer{ctrl: ctrl}
	mock.recorder = &MockPortsStreamerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortsStreamer) EXPECT() *MockPortsStreamerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockPortsStreamer) Start(ctx context.Context, r io.Reader) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", ctx, r)
}

// Start indicates an expected call of Start.
func (mr *MockPortsStreamerMockRecorder) Start(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPortsStreamer)(nil).Start), ctx, r)
}

// Stream mocks base method.
func (m *MockPortsStreamer) Stream() chan portsstream.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream")
	ret0, _ := ret[0].(chan portsstream.Result)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockPortsStreamerMockRecorder) Stream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockPortsStreamer)(nil).Stream))
}
