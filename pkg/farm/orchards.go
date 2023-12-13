package farm

import "handlers/pkg/common"

type (
	OrhcadFarm interface {
		common.RunHandler
		PickApple(string) string
		PickMango(string) string
	}

	OrchadFarmO struct {
		common.RunHandler
		desc string
	}

	orchadFarmKey  struct{ common.HandlerKey }
	orchadFarmInit struct{}
)

func (k orchadFarmKey) String() string { return "orchad farm key" }

var (
	OrchadFarmKey = orchadFarmKey{}
)

func NewOrchadFarm(args ...interface{}) OrhcadFarm {
	return &OrchadFarmO{
		desc: "orchad farm instance",
	}
}

func (p OrchadFarmO) String() string                 { return p.desc }
func (p *OrchadFarmO) Get() interface{}              { return p }
func (p *OrchadFarmO) PickApple(which string) string { return which + " apple picked" }
func (p *OrchadFarmO) PickMango(which string) string { return which + " mango picked" }

func GetOrchadFarmFromCommon() (OrhcadFarm, bool) {
	p, ok := common.GetHandler(OrchadFarmKey)
	if !ok {
		return nil, ok
	}
	return p.(OrhcadFarm), ok
}

func MustGetOrchadFarmFromCommon() OrhcadFarm {
	return common.MustGetHandler(OrchadFarmKey).(OrhcadFarm)
}

func (a orchadFarmInit) Init(args ...interface{}) common.RunHandler { return NewOrchadFarm(args) }

func init() {
	common.AddInitHandler(OrchadFarmKey, orchadFarmInit{})
}
