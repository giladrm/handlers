package store

import "handlers/pkg/common"

type (
	AppleStore interface {
		common.RunHandler
		GetSmith() string
		GetPinkLady() string
	}

	AppleStoreO struct {
		common.RunHandler
		desc string
	}

	appleKey       struct{ common.HandlerKey }
	appleStoreInit struct{}
)

func (k appleKey) String() string { return "apple store key" }

var (
	AppleStoreKey = appleKey{}
)

func NewAppleStore(args ...interface{}) AppleStore {
	return &AppleStoreO{
		desc: "Apple Store instance",
	}
}

func (p AppleStoreO) String() string       { return p.desc }
func (p *AppleStoreO) Get() interface{}    { return p }
func (p *AppleStoreO) GetSmith() string    { return "smith apple" }
func (p *AppleStoreO) GetPinkLady() string { return "pink lady apple" }

func GetAppleStoreFromCommon() (AppleStore, bool) {
	p, ok := common.GetHandler(AppleStoreKey)
	if !ok {
		return nil, ok
	}
	return p.(AppleStore), ok
}

func MustGetAppleStoreFromCommon() AppleStore {
	return common.MustGetHandler(AppleStoreKey).(AppleStore)
}

func (a appleStoreInit) Init(args ...interface{}) common.RunHandler { return NewAppleStore(args) }

func init() {
	common.AddInitHandler(AppleStoreKey, appleStoreInit{})
}
