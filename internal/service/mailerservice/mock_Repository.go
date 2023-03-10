// Code generated by mockery v2.15.0. DO NOT EDIT.

package mailerservice

import (
	context "context"

	mailgun "github.com/mailgun/mailgun-go/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// NewMessage provides a mock function with given fields: sender, subject, recipient
func (_m *MockRepository) NewMessage(sender string, subject string, recipient string) *mailgun.Message {
	ret := _m.Called(sender, subject, recipient)

	var r0 *mailgun.Message
	if rf, ok := ret.Get(0).(func(string, string, string) *mailgun.Message); ok {
		r0 = rf(sender, subject, recipient)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mailgun.Message)
		}
	}

	return r0
}

// SendMail provides a mock function with given fields: ctx, messages
func (_m *MockRepository) SendMail(ctx context.Context, messages *mailgun.Message) (string, error) {
	ret := _m.Called(ctx, messages)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *mailgun.Message) string); ok {
		r0 = rf(ctx, messages)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *mailgun.Message) error); ok {
		r1 = rf(ctx, messages)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
