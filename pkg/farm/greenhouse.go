package farm

import "handlers/pkg/common"

type (
	GreenHouseFarm interface {
		common.RunHandler
		PickTomato(string) string
		PickLettuce(string) string
	}

	GreenHouseFarmO struct {
		desc string
	}
)

func NewGreenHouseFarm(args ...interface{}) GreenHouseFarm {
	return &GreenHouseFarmO{
		desc: "greenhouse farm instance",
	}
}

func (p *GreenHouseFarmO) String() string                  { return p.desc }
func (p *GreenHouseFarmO) Get() interface{}                { return p }
func (p *GreenHouseFarmO) PickTomato(which string) string  { return which + " tomato picked" }
func (p *GreenHouseFarmO) PickLettuce(which string) string { return which + " lettuce picked" }

// register as a handler and implement retrieval func

func greenHouseFarmKeyString() string                           { return "greenhouse farm key" }
func initGreenHousedFarm(args ...interface{}) common.RunHandler { return NewOrchadFarm(args) }

var (
	GreenHouseFarmKey  = farmKey{keyString: greenHouseFarmKeyString()}
	greenHouseFarmInit = farmInit{initFunc: initGreenHousedFarm}
)

func GetGreenHouseFarmFromCommon() (GreenHouseFarm, bool) {
	p, ok := common.GetHandler(GreenHouseFarmKey)
	if !ok {
		return nil, ok
	}
	r, ok := p.(GreenHouseFarm)
	return r, ok
}

func MustGetGreenHouseFarmFromCommon() GreenHouseFarm {
	return common.MustGetHandler(GreenHouseFarmKey).(GreenHouseFarm)
}

func init() {
	common.AddInitHandler(GreenHouseFarmKey, greenHouseFarmInit, 1)
}
