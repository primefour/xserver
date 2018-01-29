// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/primefour/xserver/model"
import store "github.com/primefour/xserver/store"

// TokenStore is an autogenerated mock type for the TokenStore type
type TokenStore struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields:
func (_m *TokenStore) Cleanup() {
	_m.Called()
}

// Delete provides a mock function with given fields: token
func (_m *TokenStore) Delete(token string) store.StoreChannel {
	ret := _m.Called(token)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByToken provides a mock function with given fields: token
func (_m *TokenStore) GetByToken(token string) store.StoreChannel {
	ret := _m.Called(token)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: recovery
func (_m *TokenStore) Save(recovery *model.Token) store.StoreChannel {
	ret := _m.Called(recovery)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Token) store.StoreChannel); ok {
		r0 = rf(recovery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
