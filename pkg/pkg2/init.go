package pkg2

import (
	"handlers/pkg/common"
)

type (
	pkg2Init struct{}

	pkg2Key struct {
		common.HandlerKey
	}
)

func (k pkg2Key) String() string {
	return "pkg2 key"
}

var (
	Pkg2K = pkg2Key{}
)

func (p pkg2Init) Init(args ...interface{}) common.RunHandler {
	return NewPkg2(args)
}

func init() {
	common.AddInitHandler(Pkg2K, pkg2Init{})
}
