// Code generated by mockery v2.15.0. DO NOT EDIT.

package tokenservice

import (
	context "context"

	ent "github.com/naofel1/api-golang-template/internal/ent"
	mock "github.com/stretchr/testify/mock"

	primitive "github.com/naofel1/api-golang-template/internal/primitive"

	uuid "github.com/google/uuid"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

// GetRoleFromIDToken provides a mock function with given fields: tokenString
func (_m *MockInterface) GetRoleFromIDToken(tokenString string) (primitive.Roles, error) {
	ret := _m.Called(tokenString)

	var r0 primitive.Roles
	if rf, ok := ret.Get(0).(func(string) primitive.Roles); ok {
		r0 = rf(tokenString)
	} else {
		r0 = ret.Get(0).(primitive.Roles)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPairFromAdmin provides a mock function with given fields: ctx, u, prevTokenID
func (_m *MockInterface) NewPairFromAdmin(ctx context.Context, u *ent.Admin, prevTokenID string) (*PairToken, error) {
	ret := _m.Called(ctx, u, prevTokenID)

	var r0 *PairToken
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Admin, string) *PairToken); ok {
		r0 = rf(ctx, u, prevTokenID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PairToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Admin, string) error); ok {
		r1 = rf(ctx, u, prevTokenID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPairFromStudent provides a mock function with given fields: ctx, u, prevTokenID
func (_m *MockInterface) NewPairFromStudent(ctx context.Context, u *ent.Student, prevTokenID string) (*PairToken, error) {
	ret := _m.Called(ctx, u, prevTokenID)

	var r0 *PairToken
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Student, string) *PairToken); ok {
		r0 = rf(ctx, u, prevTokenID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PairToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Student, string) error); ok {
		r1 = rf(ctx, u, prevTokenID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Signout provides a mock function with given fields: ctx, uid
func (_m *MockInterface) Signout(ctx context.Context, uid uuid.UUID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateAdminIDToken provides a mock function with given fields: ctx, tokenString
func (_m *MockInterface) ValidateAdminIDToken(ctx context.Context, tokenString string) (*ent.Admin, error) {
	ret := _m.Called(ctx, tokenString)

	var r0 *ent.Admin
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.Admin); ok {
		r0 = rf(ctx, tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Admin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateRefreshToken provides a mock function with given fields: ctx, refreshTokenString
func (_m *MockInterface) ValidateRefreshToken(ctx context.Context, refreshTokenString string) (*RefreshToken, error) {
	ret := _m.Called(ctx, refreshTokenString)

	var r0 *RefreshToken
	if rf, ok := ret.Get(0).(func(context.Context, string) *RefreshToken); ok {
		r0 = rf(ctx, refreshTokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RefreshToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, refreshTokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateStudentIDToken provides a mock function with given fields: ctx, tokenString
func (_m *MockInterface) ValidateStudentIDToken(ctx context.Context, tokenString string) (*ent.Student, error) {
	ret := _m.Called(ctx, tokenString)

	var r0 *ent.Student
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.Student); ok {
		r0 = rf(ctx, tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterface(t mockConstructorTestingTNewMockInterface) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}