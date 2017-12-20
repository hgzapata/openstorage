// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/volume (interfaces: VolumeDriver)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	reflect "reflect"
)

// MockVolumeDriver is a mock of VolumeDriver interface
type MockVolumeDriver struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeDriverMockRecorder
}

// MockVolumeDriverMockRecorder is the mock recorder for MockVolumeDriver
type MockVolumeDriverMockRecorder struct {
	mock *MockVolumeDriver
}

// NewMockVolumeDriver creates a new mock instance
func NewMockVolumeDriver(ctrl *gomock.Controller) *MockVolumeDriver {
	mock := &MockVolumeDriver{ctrl: ctrl}
	mock.recorder = &MockVolumeDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVolumeDriver) EXPECT() *MockVolumeDriverMockRecorder {
	return m.recorder
}

// Attach mocks base method
func (m *MockVolumeDriver) Attach(arg0 string, arg1 map[string]string) (string, error) {
	ret := m.ctrl.Call(m, "Attach", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attach indicates an expected call of Attach
func (mr *MockVolumeDriverMockRecorder) Attach(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockVolumeDriver)(nil).Attach), arg0, arg1)
}

// Create mocks base method
func (m *MockVolumeDriver) Create(arg0 *api.VolumeLocator, arg1 *api.Source, arg2 *api.VolumeSpec) (string, error) {
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockVolumeDriverMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVolumeDriver)(nil).Create), arg0, arg1, arg2)
}

// CredsCreate mocks base method
func (m *MockVolumeDriver) CredsCreate(arg0 map[string]string) (string, error) {
	ret := m.ctrl.Call(m, "CredsCreate", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredsCreate indicates an expected call of CredsCreate
func (mr *MockVolumeDriverMockRecorder) CredsCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsCreate", reflect.TypeOf((*MockVolumeDriver)(nil).CredsCreate), arg0)
}

// CredsDelete mocks base method
func (m *MockVolumeDriver) CredsDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "CredsDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CredsDelete indicates an expected call of CredsDelete
func (mr *MockVolumeDriverMockRecorder) CredsDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsDelete", reflect.TypeOf((*MockVolumeDriver)(nil).CredsDelete), arg0)
}

// CredsEnumerate mocks base method
func (m *MockVolumeDriver) CredsEnumerate() (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "CredsEnumerate")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredsEnumerate indicates an expected call of CredsEnumerate
func (mr *MockVolumeDriverMockRecorder) CredsEnumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsEnumerate", reflect.TypeOf((*MockVolumeDriver)(nil).CredsEnumerate))
}

// CredsValidate mocks base method
func (m *MockVolumeDriver) CredsValidate(arg0 string) error {
	ret := m.ctrl.Call(m, "CredsValidate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CredsValidate indicates an expected call of CredsValidate
func (mr *MockVolumeDriverMockRecorder) CredsValidate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsValidate", reflect.TypeOf((*MockVolumeDriver)(nil).CredsValidate), arg0)
}

// Delete mocks base method
func (m *MockVolumeDriver) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockVolumeDriverMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVolumeDriver)(nil).Delete), arg0)
}

// Detach mocks base method
func (m *MockVolumeDriver) Detach(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "Detach", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Detach indicates an expected call of Detach
func (mr *MockVolumeDriverMockRecorder) Detach(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detach", reflect.TypeOf((*MockVolumeDriver)(nil).Detach), arg0, arg1)
}

// Enumerate mocks base method
func (m *MockVolumeDriver) Enumerate(arg0 *api.VolumeLocator, arg1 map[string]string) ([]*api.Volume, error) {
	ret := m.ctrl.Call(m, "Enumerate", arg0, arg1)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockVolumeDriverMockRecorder) Enumerate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockVolumeDriver)(nil).Enumerate), arg0, arg1)
}

// Flush mocks base method
func (m *MockVolumeDriver) Flush(arg0 string) error {
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockVolumeDriverMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockVolumeDriver)(nil).Flush), arg0)
}

// GetActiveRequests mocks base method
func (m *MockVolumeDriver) GetActiveRequests() (*api.ActiveRequests, error) {
	ret := m.ctrl.Call(m, "GetActiveRequests")
	ret0, _ := ret[0].(*api.ActiveRequests)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveRequests indicates an expected call of GetActiveRequests
func (mr *MockVolumeDriverMockRecorder) GetActiveRequests() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveRequests", reflect.TypeOf((*MockVolumeDriver)(nil).GetActiveRequests))
}

// Inspect mocks base method
func (m *MockVolumeDriver) Inspect(arg0 []string) ([]*api.Volume, error) {
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockVolumeDriverMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockVolumeDriver)(nil).Inspect), arg0)
}

// Mount mocks base method
func (m *MockVolumeDriver) Mount(arg0, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "Mount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockVolumeDriverMockRecorder) Mount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockVolumeDriver)(nil).Mount), arg0, arg1, arg2)
}

// MountedAt mocks base method
func (m *MockVolumeDriver) MountedAt(arg0 string) string {
	ret := m.ctrl.Call(m, "MountedAt", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// MountedAt indicates an expected call of MountedAt
func (mr *MockVolumeDriverMockRecorder) MountedAt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountedAt", reflect.TypeOf((*MockVolumeDriver)(nil).MountedAt), arg0)
}

// Name mocks base method
func (m *MockVolumeDriver) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockVolumeDriverMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockVolumeDriver)(nil).Name))
}

// Quiesce mocks base method
func (m *MockVolumeDriver) Quiesce(arg0 string, arg1 uint64, arg2 string) error {
	ret := m.ctrl.Call(m, "Quiesce", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Quiesce indicates an expected call of Quiesce
func (mr *MockVolumeDriverMockRecorder) Quiesce(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quiesce", reflect.TypeOf((*MockVolumeDriver)(nil).Quiesce), arg0, arg1, arg2)
}

// Read mocks base method
func (m *MockVolumeDriver) Read(arg0 string, arg1 []byte, arg2 uint64, arg3 int64) (int64, error) {
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockVolumeDriverMockRecorder) Read(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockVolumeDriver)(nil).Read), arg0, arg1, arg2, arg3)
}

// Restore mocks base method
func (m *MockVolumeDriver) Restore(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockVolumeDriverMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockVolumeDriver)(nil).Restore), arg0, arg1)
}

// Set mocks base method
func (m *MockVolumeDriver) Set(arg0 string, arg1 *api.VolumeLocator, arg2 *api.VolumeSpec) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockVolumeDriverMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockVolumeDriver)(nil).Set), arg0, arg1, arg2)
}

// Shutdown mocks base method
func (m *MockVolumeDriver) Shutdown() {
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockVolumeDriverMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockVolumeDriver)(nil).Shutdown))
}

// SnapEnumerate mocks base method
func (m *MockVolumeDriver) SnapEnumerate(arg0 []string, arg1 map[string]string) ([]*api.Volume, error) {
	ret := m.ctrl.Call(m, "SnapEnumerate", arg0, arg1)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapEnumerate indicates an expected call of SnapEnumerate
func (mr *MockVolumeDriverMockRecorder) SnapEnumerate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapEnumerate", reflect.TypeOf((*MockVolumeDriver)(nil).SnapEnumerate), arg0, arg1)
}

// Snapshot mocks base method
func (m *MockVolumeDriver) Snapshot(arg0 string, arg1 bool, arg2 *api.VolumeLocator) (string, error) {
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockVolumeDriverMockRecorder) Snapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockVolumeDriver)(nil).Snapshot), arg0, arg1, arg2)
}

// Stats mocks base method
func (m *MockVolumeDriver) Stats(arg0 string, arg1 bool) (*api.Stats, error) {
	ret := m.ctrl.Call(m, "Stats", arg0, arg1)
	ret0, _ := ret[0].(*api.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockVolumeDriverMockRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockVolumeDriver)(nil).Stats), arg0, arg1)
}

// Status mocks base method
func (m *MockVolumeDriver) Status() [][2]string {
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].([][2]string)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockVolumeDriverMockRecorder) Status() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockVolumeDriver)(nil).Status))
}

// Type mocks base method
func (m *MockVolumeDriver) Type() api.DriverType {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(api.DriverType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockVolumeDriverMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockVolumeDriver)(nil).Type))
}

// Unmount mocks base method
func (m *MockVolumeDriver) Unmount(arg0, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "Unmount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockVolumeDriverMockRecorder) Unmount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockVolumeDriver)(nil).Unmount), arg0, arg1, arg2)
}

// Unquiesce mocks base method
func (m *MockVolumeDriver) Unquiesce(arg0 string) error {
	ret := m.ctrl.Call(m, "Unquiesce", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unquiesce indicates an expected call of Unquiesce
func (mr *MockVolumeDriverMockRecorder) Unquiesce(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unquiesce", reflect.TypeOf((*MockVolumeDriver)(nil).Unquiesce), arg0)
}

// UsedSize mocks base method
func (m *MockVolumeDriver) UsedSize(arg0 string) (uint64, error) {
	ret := m.ctrl.Call(m, "UsedSize", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsedSize indicates an expected call of UsedSize
func (mr *MockVolumeDriverMockRecorder) UsedSize(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsedSize", reflect.TypeOf((*MockVolumeDriver)(nil).UsedSize), arg0)
}

// Write mocks base method
func (m *MockVolumeDriver) Write(arg0 string, arg1 []byte, arg2 uint64, arg3 int64) (int64, error) {
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockVolumeDriverMockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockVolumeDriver)(nil).Write), arg0, arg1, arg2, arg3)
}
