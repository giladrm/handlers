package distribution_test

import (
	"handlers/pkg/common"
	"handlers/pkg/distribution"

	"github.com/stretchr/testify/mock"
)

type (
	TruckDistributionMock struct {
		common.RunHandler
		mock.Mock
	}

	truckDistributionMockKey  struct{ common.HandlerKey }
	truckDistributionMockInit struct{}
)

func (k truckDistributionMockKey) String() string { return "truck distribution mock key" }

var (
	TruckDistributionMockKey = truckDistributionMockKey{}
)

func NewTruckDistributionMock(args ...interface{}) distribution.TruckDistribution {
	return &TruckDistributionMock{}
}

func (m *TruckDistributionMock) String() string             { return m.Called().String(0) }
func (m *TruckDistributionMock) Get() interface{}           { return m.Called().Get(0) }
func (m *TruckDistributionMock) Load(which string) string   { return m.Called(which).Get(0).(string) }
func (m *TruckDistributionMock) Unload(which string) string { return m.Called(which).Get(0).(string) }

func GetTruckDistributionMockFromCommon() (distribution.TruckDistribution, bool) {
	p, ok := common.GetHandler(TruckDistributionMockKey)
	if !ok {
		return nil, ok
	}
	return p.(distribution.TruckDistribution), ok
}

func MustGetTruckDistributionMockFromCommon() distribution.TruckDistribution {
	return common.MustGetHandler(TruckDistributionMockKey).(distribution.TruckDistribution)
}

func (a truckDistributionMockInit) Init(args ...interface{}) common.RunHandler {
	return NewTruckDistributionMock(args)
}

func init() {
	common.AddInitHandler(TruckDistributionMockKey, truckDistributionMockInit{})
}
