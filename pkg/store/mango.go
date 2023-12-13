package store

import "handlers/pkg/common"

type (
	MangoStore interface {
		common.RunHandler
		GetMaya() string
		GetIrwin() string
	}

	MangoStoreO struct {
		common.RunHandler
		desc string
	}

	mangoKey       struct{ common.HandlerKey }
	mangoStoreInit struct{}
)

func (k mangoKey) String() string { return "mango store key" }

var (
	MangoStoreKey = mangoKey{}
)

func NewMangoStore(args ...interface{}) MangoStore {
	return &MangoStoreO{
		desc: "Mango Store instance",
	}
}

func (p *MangoStoreO) String() string   { return p.desc }
func (p *MangoStoreO) Get() interface{} { return p }
func (p *MangoStoreO) GetMaya() string  { return "maya mango" }
func (p *MangoStoreO) GetIrwin() string { return "irwin mango" }

func GetMangoStoreFromCommon() (MangoStore, bool) {
	p, ok := common.GetHandler(MangoStoreKey)
	if !ok {
		return nil, ok
	}
	return p.(MangoStore), ok
}

func MustGetMangoStoreFromCommon() MangoStore {
	return common.MustGetHandler(MangoStoreKey).(MangoStore)
}

func (a mangoStoreInit) Init(args ...interface{}) common.RunHandler { return NewMangoStore(args) }

func init() {
	common.AddInitHandler(MangoStoreKey, mangoStoreInit{})
}
