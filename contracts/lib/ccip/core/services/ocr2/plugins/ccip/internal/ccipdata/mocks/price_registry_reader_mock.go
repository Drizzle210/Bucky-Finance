// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	ccipdata "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"

	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PriceRegistryReader is an autogenerated mock type for the PriceRegistryReader type
type PriceRegistryReader struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *PriceRegistryReader) Address() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *PriceRegistryReader) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFeeTokens provides a mock function with given fields: ctx
func (_m *PriceRegistryReader) GetFeeTokens(ctx context.Context) ([]common.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFeeTokens")
	}

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]common.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []common.Address); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPriceUpdatesCreatedAfter provides a mock function with given fields: ctx, chainSelector, ts, confs
func (_m *PriceRegistryReader) GetGasPriceUpdatesCreatedAfter(ctx context.Context, chainSelector uint64, ts time.Time, confs int) ([]ccipdata.Event[ccipdata.GasPriceUpdate], error) {
	ret := _m.Called(ctx, chainSelector, ts, confs)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPriceUpdatesCreatedAfter")
	}

	var r0 []ccipdata.Event[ccipdata.GasPriceUpdate]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, time.Time, int) ([]ccipdata.Event[ccipdata.GasPriceUpdate], error)); ok {
		return rf(ctx, chainSelector, ts, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, time.Time, int) []ccipdata.Event[ccipdata.GasPriceUpdate]); ok {
		r0 = rf(ctx, chainSelector, ts, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipdata.Event[ccipdata.GasPriceUpdate])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, time.Time, int) error); ok {
		r1 = rf(ctx, chainSelector, ts, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenPriceUpdatesCreatedAfter provides a mock function with given fields: ctx, ts, confs
func (_m *PriceRegistryReader) GetTokenPriceUpdatesCreatedAfter(ctx context.Context, ts time.Time, confs int) ([]ccipdata.Event[ccipdata.TokenPriceUpdate], error) {
	ret := _m.Called(ctx, ts, confs)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenPriceUpdatesCreatedAfter")
	}

	var r0 []ccipdata.Event[ccipdata.TokenPriceUpdate]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, int) ([]ccipdata.Event[ccipdata.TokenPriceUpdate], error)); ok {
		return rf(ctx, ts, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, int) []ccipdata.Event[ccipdata.TokenPriceUpdate]); ok {
		r0 = rf(ctx, ts, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipdata.Event[ccipdata.TokenPriceUpdate])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, int) error); ok {
		r1 = rf(ctx, ts, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenPrices provides a mock function with given fields: ctx, wantedTokens
func (_m *PriceRegistryReader) GetTokenPrices(ctx context.Context, wantedTokens []common.Address) ([]ccipdata.TokenPriceUpdate, error) {
	ret := _m.Called(ctx, wantedTokens)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenPrices")
	}

	var r0 []ccipdata.TokenPriceUpdate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address) ([]ccipdata.TokenPriceUpdate, error)); ok {
		return rf(ctx, wantedTokens)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address) []ccipdata.TokenPriceUpdate); ok {
		r0 = rf(ctx, wantedTokens)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipdata.TokenPriceUpdate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []common.Address) error); ok {
		r1 = rf(ctx, wantedTokens)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokensDecimals provides a mock function with given fields: ctx, tokenAddresses
func (_m *PriceRegistryReader) GetTokensDecimals(ctx context.Context, tokenAddresses []common.Address) ([]uint8, error) {
	ret := _m.Called(ctx, tokenAddresses)

	if len(ret) == 0 {
		panic("no return value specified for GetTokensDecimals")
	}

	var r0 []uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address) ([]uint8, error)); ok {
		return rf(ctx, tokenAddresses)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address) []uint8); ok {
		r0 = rf(ctx, tokenAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint8)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []common.Address) error); ok {
		r1 = rf(ctx, tokenAddresses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPriceRegistryReader creates a new instance of PriceRegistryReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceRegistryReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceRegistryReader {
	mock := &PriceRegistryReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}