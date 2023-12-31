// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "github.com/Zhenya671/golang-test-task/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// PayOff provides a mock function with given fields: userId, input
func (_m *IUserService) PayOff(userId string, input model.Debt) (model.Debt, error) {
	ret := _m.Called(userId, input)

	var r0 model.Debt
	var r1 error
	if rf, ok := ret.Get(0).(func(string, model.Debt) (model.Debt, error)); ok {
		return rf(userId, input)
	}
	if rf, ok := ret.Get(0).(func(string, model.Debt) model.Debt); ok {
		r0 = rf(userId, input)
	} else {
		r0 = ret.Get(0).(model.Debt)
	}

	if rf, ok := ret.Get(1).(func(string, model.Debt) error); ok {
		r1 = rf(userId, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignIn provides a mock function with given fields: logIn
func (_m *IUserService) SignIn(logIn model.User) (string, error) {
	ret := _m.Called(logIn)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (string, error)); ok {
		return rf(logIn)
	}
	if rf, ok := ret.Get(0).(func(model.User) string); ok {
		r0 = rf(logIn)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(logIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: user
func (_m *IUserService) SignUp(user model.User) (string, error) {
	ret := _m.Called(user)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(model.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SolveAlgo provides a mock function with given fields: userID, algoName, input
func (_m *IUserService) SolveAlgo(userID string, algoName string, input model.Task) (model.Task, error) {
	ret := _m.Called(userID, algoName, input)

	var r0 model.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.Task) (model.Task, error)); ok {
		return rf(userID, algoName, input)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.Task) model.Task); ok {
		r0 = rf(userID, algoName, input)
	} else {
		r0 = ret.Get(0).(model.Task)
	}

	if rf, ok := ret.Get(1).(func(string, string, model.Task) error); ok {
		r1 = rf(userID, algoName, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserService creates a new instance of IUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserService(t mockConstructorTestingTNewIUserService) *IUserService {
	mock := &IUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
