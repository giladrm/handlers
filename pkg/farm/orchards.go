package farm

import "handlers/pkg/common"

type (
	OrhcadFarm interface {
		common.RunHandler
		PickApple(string) string
		PickMango(string) string
	}

	OrchadFarmO struct {
		desc string
	}
)

func NewOrchadFarm(args ...interface{}) OrhcadFarm {
	return &OrchadFarmO{
		desc: "orchad farm instance",
	}
}

func (p *OrchadFarmO) String() string                { return p.desc }
func (p *OrchadFarmO) Get() interface{}              { return p }
func (p *OrchadFarmO) PickApple(which string) string { return which + " apple picked" }
func (p *OrchadFarmO) PickMango(which string) string { return which + " mango picked" }

// register as a handler and implement retrieval func

func orchadFarmKeyString() string                          { return "orchad farm key" }
func initOrchadFarm(args ...interface{}) common.RunHandler { return NewOrchadFarm(args) }

var (
	OrchadFarmKey  = farmKey{keyString: orchadFarmKeyString()}
	orchadFarmInit = farmInit{initFunc: initOrchadFarm}
)

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

func init() { common.AddInitHandler(OrchadFarmKey, orchadFarmInit, 1) }
