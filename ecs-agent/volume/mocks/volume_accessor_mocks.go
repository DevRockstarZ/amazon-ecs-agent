// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/volume (interfaces: TaskVolumeAccessor)

// Package mock_volume is a generated GoMock package.
package mock_volume

import (
	fs "io/fs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskVolumeAccessor is a mock of TaskVolumeAccessor interface.
type MockTaskVolumeAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockTaskVolumeAccessorMockRecorder
}

// MockTaskVolumeAccessorMockRecorder is the mock recorder for MockTaskVolumeAccessor.
type MockTaskVolumeAccessorMockRecorder struct {
	mock *MockTaskVolumeAccessor
}

// NewMockTaskVolumeAccessor creates a new mock instance.
func NewMockTaskVolumeAccessor(ctrl *gomock.Controller) *MockTaskVolumeAccessor {
	mock := &MockTaskVolumeAccessor{ctrl: ctrl}
	mock.recorder = &MockTaskVolumeAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskVolumeAccessor) EXPECT() *MockTaskVolumeAccessorMockRecorder {
	return m.recorder
}

// CopyToVolume mocks base method.
func (m *MockTaskVolumeAccessor) CopyToVolume(arg0, arg1, arg2 string, arg3 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyToVolume", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyToVolume indicates an expected call of CopyToVolume.
func (mr *MockTaskVolumeAccessorMockRecorder) CopyToVolume(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyToVolume", reflect.TypeOf((*MockTaskVolumeAccessor)(nil).CopyToVolume), arg0, arg1, arg2, arg3)
}

// DeleteAll mocks base method.
func (m *MockTaskVolumeAccessor) DeleteAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockTaskVolumeAccessorMockRecorder) DeleteAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockTaskVolumeAccessor)(nil).DeleteAll), arg0)
}