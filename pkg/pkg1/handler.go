package pkg1

import "handlers/pkg/common"

type (
	Pkg1 interface {
		common.RunHandler
		Foo() string
	}

	Pkg1Concrete struct {
		common.RunHandler
		desc string
	}
)

func NewPkg1(args ...interface{}) Pkg1 {
	return &Pkg1Concrete{
		desc: "pkg1 instance",
	}
}
func (p *Pkg1Concrete) Get() interface{} {
	return p
}
func (p Pkg1Concrete) String() string {
	return p.desc
}
func (p *Pkg1Concrete) Foo() string {
	return "foo"
}
