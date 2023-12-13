package pkg2

import (
	"handlers/pkg/common"
	"handlers/pkg/pkg1"
)

type (
	Pkg2 interface {
		common.RunHandler
		Bar() string
	}

	Pkg2Concrete struct {
		common.RunHandler
		desc string
		pkg1 pkg1.Pkg1
	}
)

func NewPkg2(args ...interface{}) Pkg2 {
	pkg1 := common.MustGetHandler(pkg1.Pkg1K).(pkg1.Pkg1)
	return &Pkg2Concrete{pkg1: pkg1,
		desc: "pkg2 isntance (with 1 inside)"}
}
func (p *Pkg2Concrete) Get() interface{} {
	return p
}
func (p Pkg2Concrete) String() string {
	return p.desc + ":" + p.pkg1.String()
}
func (p *Pkg2Concrete) Bar() string {
	return "bar" + ":" + p.pkg1.Foo()
}

func MustGetPkg2FromCommon() Pkg2 {
	return common.MustGetHandler(Pkg2K).(Pkg2)
}
