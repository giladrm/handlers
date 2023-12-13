package distribution_test

import (
	"handlers/pkg/common"
	"handlers/pkg/distribution"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruck_flow(t *testing.T) {
	common.InitSome([]common.HandlerKey{distribution.TruckDistributionMockKey})

	mockTruck := distribution.MustGetTruckDistributionMockFromCommon()
	assert.NotNil(t, mockTruck)
	a := mockTruck.(*distribution.TruckDistributionMock)
	a.On("Load", "abc").Return("abc")

	mockTruck.Load("abc")

	a.AssertCalled(t, "Load", "abc")
}
