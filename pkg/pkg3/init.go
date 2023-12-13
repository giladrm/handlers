package pkg3

import "handlers/pkg/common"

type (
	pkg3Init struct {
	}

	pkg3Key struct {
		common.HandlerKey
	}
)

func (k pkg3Key) String() string {
	return "pkg3 key"
}

var (
	Pkg3K = pkg3Key{}
)

func (p pkg3Init) Init(args ...interface{}) common.RunHandler {
	return NewPkg3(args)
}

func init() {
	common.AddInitHandler(Pkg3K, pkg3Init{})
}
