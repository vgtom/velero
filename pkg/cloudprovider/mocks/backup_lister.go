/*
Copyright 2018 the Heptio Ark contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "github.com/heptio/ark/pkg/apis/ark/v1"

// BackupLister is an autogenerated mock type for the BackupLister type
type BackupLister struct {
	mock.Mock
}

// ListBackups provides a mock function with given fields: bucket
func (_m *BackupLister) ListBackups(bucket string) ([]*v1.Backup, error) {
	ret := _m.Called(bucket)

	var r0 []*v1.Backup
	if rf, ok := ret.Get(0).(func(string) []*v1.Backup); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
