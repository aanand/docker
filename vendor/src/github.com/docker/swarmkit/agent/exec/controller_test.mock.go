// Automatically generated by MockGen. DO NOT EDIT!
// Source: controller.go

package exec

import (
	api "github.com/docker/swarmkit/api"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *_MockControllerRecorder
}

// Recorder for MockController (not exported)
type _MockControllerRecorder struct {
	mock *MockController
}

func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &_MockControllerRecorder{mock}
	return mock
}

func (_m *MockController) EXPECT() *_MockControllerRecorder {
	return _m.recorder
}

func (_m *MockController) Update(ctx context.Context, t *api.Task) error {
	ret := _m.ctrl.Call(_m, "Update", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1)
}

func (_m *MockController) Prepare(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Prepare", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Prepare(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0)
}

func (_m *MockController) Start(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockController) Wait(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Wait", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Wait(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Wait", arg0)
}

func (_m *MockController) Shutdown(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Shutdown", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shutdown", arg0)
}

func (_m *MockController) Terminate(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Terminate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Terminate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Terminate", arg0)
}

func (_m *MockController) Remove(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Remove", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Remove(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0)
}

func (_m *MockController) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockControllerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of ContainerStatuser interface
type MockContainerStatuser struct {
	ctrl     *gomock.Controller
	recorder *_MockContainerStatuserRecorder
}

// Recorder for MockContainerStatuser (not exported)
type _MockContainerStatuserRecorder struct {
	mock *MockContainerStatuser
}

func NewMockContainerStatuser(ctrl *gomock.Controller) *MockContainerStatuser {
	mock := &MockContainerStatuser{ctrl: ctrl}
	mock.recorder = &_MockContainerStatuserRecorder{mock}
	return mock
}

func (_m *MockContainerStatuser) EXPECT() *_MockContainerStatuserRecorder {
	return _m.recorder
}

func (_m *MockContainerStatuser) ContainerStatus(ctx context.Context) (*api.ContainerStatus, error) {
	ret := _m.ctrl.Call(_m, "ContainerStatus", ctx)
	ret0, _ := ret[0].(*api.ContainerStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContainerStatuserRecorder) ContainerStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerStatus", arg0)
}

// Mock of PortStatuser interface
type MockPortStatuser struct {
	ctrl     *gomock.Controller
	recorder *_MockPortStatuserRecorder
}

// Recorder for MockPortStatuser (not exported)
type _MockPortStatuserRecorder struct {
	mock *MockPortStatuser
}

func NewMockPortStatuser(ctrl *gomock.Controller) *MockPortStatuser {
	mock := &MockPortStatuser{ctrl: ctrl}
	mock.recorder = &_MockPortStatuserRecorder{mock}
	return mock
}

func (_m *MockPortStatuser) EXPECT() *_MockPortStatuserRecorder {
	return _m.recorder
}

func (_m *MockPortStatuser) PortStatus(ctx context.Context) (*api.PortStatus, error) {
	ret := _m.ctrl.Call(_m, "PortStatus", ctx)
	ret0, _ := ret[0].(*api.PortStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPortStatuserRecorder) PortStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PortStatus", arg0)
}
