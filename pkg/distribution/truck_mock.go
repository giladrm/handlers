package distribution

import (
	"handlers/pkg/common"

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

func NewTruckDistributionMock(args ...interface{}) TruckDistribution {
	return &TruckDistributionMock{}
}

func (m *TruckDistributionMock) String() string             { return m.Called().Get(0).(string) }
func (m *TruckDistributionMock) Get() interface{}           { return m.Called().Get(0) }
func (m *TruckDistributionMock) Load(which string) string   { return m.Called(which).Get(0).(string) }
func (m *TruckDistributionMock) Unload(which string) string { return m.Called(which).Get(0).(string) }

func GetTruckDistributionMockFromCommon() (TruckDistribution, bool) {
	p, ok := common.GetHandler(TruckDistributionMockKey)
	if !ok {
		return nil, ok
	}
	return p.(TruckDistribution), ok
}

func MustGetTruckDistributionMockFromCommon() TruckDistribution {
	return common.MustGetHandler(TruckDistributionMockKey).(TruckDistribution)
}

func (a truckDistributionMockInit) Init(args ...interface{}) common.RunHandler {
	return NewTruckDistributionMock(args)
}

func init() {
	common.AddInitHandler(TruckDistributionMockKey, truckDistributionMockInit{})
}
