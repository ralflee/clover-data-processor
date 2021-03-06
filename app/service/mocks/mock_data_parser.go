// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	model "clover-data-processor/app/model"

	mock "github.com/stretchr/testify/mock"
)

// DataParser is an autogenerated mock type for the DataParser type
type DataParser struct {
	mock.Mock
}

// ConstructRecords provides a mock function with given fields: dataFilePath, spec
func (_m *DataParser) ConstructRecords(dataFilePath string, spec *model.Spec) ([]*model.Record, error) {
	ret := _m.Called(dataFilePath, spec)

	var r0 []*model.Record
	if rf, ok := ret.Get(0).(func(string, *model.Spec) []*model.Record); ok {
		r0 = rf(dataFilePath, spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Record)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *model.Spec) error); ok {
		r1 = rf(dataFilePath, spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConstructSpec provides a mock function with given fields: specFilePath
func (_m *DataParser) ConstructSpec(specFilePath string) (*model.Spec, error) {
	ret := _m.Called(specFilePath)

	var r0 *model.Spec
	if rf, ok := ret.Get(0).(func(string) *model.Spec); ok {
		r0 = rf(specFilePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Spec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(specFilePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
