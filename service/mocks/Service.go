// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	models "elibrary/models"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CheckAvailability provides a mock function with given fields: title
func (_m *Service) CheckAvailability(title string) (*models.BookDetail, error) {
	ret := _m.Called(title)

	if len(ret) == 0 {
		panic("no return value specified for CheckAvailability")
	}

	var r0 *models.BookDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.BookDetail, error)); ok {
		return rf(title)
	}
	if rf, ok := ret.Get(0).(func(string) *models.BookDetail); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BookDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
