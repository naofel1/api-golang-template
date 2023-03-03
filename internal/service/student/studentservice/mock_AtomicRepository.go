// Code generated by mockery v2.15.0. DO NOT EDIT.

package studentservice

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAtomicRepository is an autogenerated mock type for the AtomicRepository type
type MockAtomicRepository struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockAtomicRepository) Execute(_a0 context.Context, _a1 AtomicOperation) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AtomicOperation) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockAtomicRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAtomicRepository creates a new instance of MockAtomicRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAtomicRepository(t mockConstructorTestingTNewMockAtomicRepository) *MockAtomicRepository {
	mock := &MockAtomicRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}