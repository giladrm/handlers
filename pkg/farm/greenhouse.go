package farm

import "handlers/pkg/common"

type (
	GreenHouseFarm interface {
		common.RunHandler
		PickTomato(string) string
		PickLettuce(string) string
	}

	GreenHouseFarmO struct {
		common.RunHandler
		desc string
	}

	greenHouseFarmKey  struct{ common.HandlerKey }
	greenHouseFarmInit struct{}
)

func (k greenHouseFarmKey) String() string { return "greenhouse farm key" }

var (
	GreenHouseFarmKey = greenHouseFarmKey{}
)

func NewGreenHouseFarm(args ...interface{}) GreenHouseFarm {
	return &GreenHouseFarmO{
		desc: "orchad farm instance",
	}
}

func (p GreenHouseFarmO) String() string                   { return p.desc }
func (p *GreenHouseFarmO) Get() interface{}                { return p }
func (p *GreenHouseFarmO) PickTomato(which string) string  { return which + " tomato picked" }
func (p *GreenHouseFarmO) PickLettuce(which string) string { return which + " lettuce picked" }

func GetGreenHouseFarmFromCommon() (GreenHouseFarm, bool) {
	p, ok := common.GetHandler(GreenHouseFarmKey)
	if !ok {
		return nil, ok
	}
	return p.(GreenHouseFarm), ok
}

func MustGetGreenHuseFarmFromCommon() GreenHouseFarm {
	return common.MustGetHandler(GreenHouseFarmKey).(GreenHouseFarm)
}

func (a greenHouseFarmInit) Init(args ...interface{}) common.RunHandler {
	return NewGreenHouseFarm(args)
}

func init() {
	common.AddInitHandler(GreenHouseFarmKey, greenHouseFarmInit{})
}
