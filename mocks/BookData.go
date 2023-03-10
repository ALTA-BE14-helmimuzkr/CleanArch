// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	book "api/features/book"

	mock "github.com/stretchr/testify/mock"
)

// BookData is an autogenerated mock type for the BookData type
type BookData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, newBook
func (_m *BookData) Add(userID int, newBook book.Core) (book.Core, error) {
	ret := _m.Called(userID, newBook)

	var r0 book.Core
	if rf, ok := ret.Get(0).(func(int, book.Core) book.Core); ok {
		r0 = rf(userID, newBook)
	} else {
		r0 = ret.Get(0).(book.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, book.Core) error); ok {
		r1 = rf(userID, newBook)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, bookID
func (_m *BookData) Delete(userID int, bookID int) error {
	ret := _m.Called(userID, bookID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(userID, bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBook provides a mock function with given fields:
func (_m *BookData) GetAllBook() ([]book.Core, error) {
	ret := _m.Called()

	var r0 []book.Core
	if rf, ok := ret.Get(0).(func() []book.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]book.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyBook provides a mock function with given fields: userID
func (_m *BookData) MyBook(userID int) ([]book.Core, error) {
	ret := _m.Called(userID)

	var r0 []book.Core
	if rf, ok := ret.Get(0).(func(int) []book.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]book.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userID, bookID, updatedData
func (_m *BookData) Update(userID int, bookID int, updatedData book.Core) (book.Core, error) {
	ret := _m.Called(userID, bookID, updatedData)

	var r0 book.Core
	if rf, ok := ret.Get(0).(func(int, int, book.Core) book.Core); ok {
		r0 = rf(userID, bookID, updatedData)
	} else {
		r0 = ret.Get(0).(book.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, book.Core) error); ok {
		r1 = rf(userID, bookID, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookData interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookData creates a new instance of BookData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookData(t mockConstructorTestingTNewBookData) *BookData {
	mock := &BookData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
