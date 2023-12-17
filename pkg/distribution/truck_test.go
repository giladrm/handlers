package distribution_test

import (
	"handlers/pkg/common"
	"handlers/pkg/distribution"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruck_mock_flow(t *testing.T) {
	defer common.ResetHandlers()
	common.InitSome([]common.HandlerKey{TruckDistributionMockKey})

	mockTruck_, ok := GetTruckDistributionMockFromCommon()
	assert.True(t, ok)
	assert.NotNil(t, mockTruck_)

	mockTruck := MustGetTruckDistributionMockFromCommon()
	assert.NotNil(t, mockTruck)
	assert.Equal(t, mockTruck, mockTruck_)

	a := mockTruck.(*TruckDistributionMock)

	a.On("Load", "abc").Return("abc")
	mockTruck.Load("abc")
	a.AssertCalled(t, "Load", "abc")

	a.On("Unload", "xyz").Return("xyz")
	mockTruck.Unload("xyz")
	a.AssertCalled(t, "Unload", "xyz")

	stringResult := "StringResult"
	a.On("String").Return(stringResult)
	assert.Equal(t, stringResult, mockTruck.String())
	a.AssertCalled(t, "String")

	dummy := struct{}{}
	a.On("Get").Return(dummy)
	assert.Equal(t, mockTruck.Get(), dummy)
	a.AssertCalled(t, "Get")
}

func TestTruck_happy_flow(t *testing.T) {
	defer common.ResetHandlers()
	common.InitSome([]common.HandlerKey{distribution.TruckDistributionKey})

	truck := distribution.MustGetTruckDistributionFromCommon()
	assert.NotNil(t, truck)

	whatToLoad := "cargo"
	assert.Contains(t, truck.Load(whatToLoad), whatToLoad)
	assert.Contains(t, truck.Unload(whatToLoad), whatToLoad)

	assert.Contains(t, truck.String(), "instance")

	truck_, ok := distribution.GetTruckDistributionFromCommon()
	assert.True(t, ok)
	assert.NotNil(t, truck_)
	assert.Equal(t, truck_, truck.Get())

	var handler common.RunHandler
	assert.NotPanics(t, func() {
		handler = common.MustGetHandler(distribution.TruckDistributionKey)
	})
	assert.Equal(t, handler, truck.Get())
}

func TestTruck_unhappy_flow(t *testing.T) {
	defer common.ResetHandlers()
	common.InitSome([]common.HandlerKey{TruckDistributionMockKey})
	var keys []common.HandlerKey
	common.GetHandlersMap().Range(func(key common.HandlerKey, value common.RunHandler) bool { keys = append(keys, key); return true })
	assert.NotEmpty(t, keys)
	assert.NotContains(t, keys, distribution.TruckDistributionKey)

	h, ok := distribution.GetTruckDistributionFromCommon()
	assert.False(t, ok)
	assert.Nil(t, h)

	assert.Panics(t, func() {
		distribution.MustGetTruckDistributionFromCommon()
	})
}
