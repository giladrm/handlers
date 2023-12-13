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

func GetPkg1FromCommon() (Pkg1, bool) {
	p, ok := common.GetHandler(Pkg1K)
	if !ok {
		return nil, ok
	}
	return p.(Pkg1), ok
}

func MustGetPkg1FromCommon() Pkg1 {
	return common.MustGetHandler(Pkg1K).(Pkg1)
}
