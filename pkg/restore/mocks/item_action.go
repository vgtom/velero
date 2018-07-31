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
import restore "github.com/heptio/ark/pkg/restore"
import runtime "k8s.io/apimachinery/pkg/runtime"
import v1 "github.com/heptio/ark/pkg/apis/ark/v1"

// ItemAction is an autogenerated mock type for the ItemAction type
type ItemAction struct {
	mock.Mock
}

// AppliesTo provides a mock function with given fields:
func (_m *ItemAction) AppliesTo() (restore.ResourceSelector, error) {
	ret := _m.Called()

	var r0 restore.ResourceSelector
	if rf, ok := ret.Get(0).(func() restore.ResourceSelector); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(restore.ResourceSelector)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Execute provides a mock function with given fields: obj, _a1
func (_m *ItemAction) Execute(obj runtime.Unstructured, _a1 *v1.Restore) (runtime.Unstructured, error, error) {
	ret := _m.Called(obj, _a1)

	var r0 runtime.Unstructured
	if rf, ok := ret.Get(0).(func(runtime.Unstructured, *v1.Restore) runtime.Unstructured); ok {
		r0 = rf(obj, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Unstructured, *v1.Restore) error); ok {
		r1 = rf(obj, _a1)
	} else {
		r1 = ret.Error(1)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(runtime.Unstructured, *v1.Restore) error); ok {
		r2 = rf(obj, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
