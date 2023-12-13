package pkg1

import "handlers/pkg/common"

type (
	pkg1Init struct{}

	pkg1Key struct {
		common.HandlerKey
	}
)

func (k pkg1Key) String() string {
	return "pkg1 key"
}

var (
	Pkg1K = pkg1Key{}
)

// pkg1 init

func (p pkg1Init) Init(args ...interface{}) common.RunHandler {
	return NewPkg1(args)
}

func init() {
	common.AddInitHandler(Pkg1K, pkg1Init{})
}
