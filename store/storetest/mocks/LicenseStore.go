// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/primefour/xserver/model"
import store "github.com/primefour/xserver/store"

// LicenseStore is an autogenerated mock type for the LicenseStore type
type LicenseStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *LicenseStore) Get(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: license
func (_m *LicenseStore) Save(license *model.LicenseRecord) store.StoreChannel {
	ret := _m.Called(license)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.LicenseRecord) store.StoreChannel); ok {
		r0 = rf(license)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
