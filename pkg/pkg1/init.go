package pkg1

import "handlers/pkg/common"

type (
	pkg1h1Key  struct{ common.HandlerKey }
	pkg1h2Key  struct{ common.HandlerKey }
	pkg1h1Init struct{}
	pkg1h2Init struct{}
)

func (k pkg1h1Key) String() string { return "pkg1 h1 key" }

func (k pkg1h2Key) String() string { return "pkg1 h2 key" }

var (
	Pkg1h1K = pkg1h1Key{}
	Pkg1h2K = pkg1h2Key{}
)

// pkg1 init

func (p pkg1h1Init) Init(args ...interface{}) common.RunHandler { return NewPkg1Handler1(args) }
func (p pkg1h2Init) Init(args ...interface{}) common.RunHandler { return NewPkg1Handler2(args) }

func init() {
	common.AddInitHandler(Pkg1h1K, pkg1h1Init{})
	common.AddInitHandler(Pkg1h2K, pkg1h2Init{})
}
