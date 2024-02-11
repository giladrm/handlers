package distribution

import "handlers/pkg/common"

type (
	TruckDistribution interface {
		common.RunHandler
		Load(string) string
		Unload(string) string
	}

	TruckDistributionO struct {
		desc string
	}

	truckDistributionKey  struct{}
	truckDistributionInit struct{}
)

func (k truckDistributionKey) String() string { return "truck distribution key" }

var (
	TruckDistributionKey = truckDistributionKey{}
)

func NewTruckDistribution(args ...interface{}) TruckDistribution {
	return &TruckDistributionO{
		desc: "truck distribution instance",
	}
}

func (p *TruckDistributionO) String() string             { return p.desc }
func (p *TruckDistributionO) Get() interface{}           { return p }
func (p *TruckDistributionO) Load(which string) string   { return "loading " + which }
func (p *TruckDistributionO) Unload(which string) string { return "unloading " + which }

func GetTruckDistributionFromCommon() (TruckDistribution, bool) {
	p, ok := common.GetHandler(TruckDistributionKey)
	if !ok {
		return nil, ok
	}
	return p.(TruckDistribution), ok
}

func MustGetTruckDistributionFromCommon() TruckDistribution {
	return common.MustGetHandler(TruckDistributionKey).(TruckDistribution)
}

func (a truckDistributionInit) Init(args ...interface{}) common.RunHandler {
	return NewTruckDistribution(args)
}

func init() {
	common.AddInitHandler(TruckDistributionKey, truckDistributionInit{}, 1)
}
