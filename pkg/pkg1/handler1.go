package pkg1

import "handlers/pkg/common"

type (
	Pkg1Handler1 interface {
		common.RunHandler
		Foo1() string
	}

	Pkg1Handler1Concrete struct {
		common.RunHandler
		desc string
	}
)

func NewPkg1Handler1(args ...interface{}) Pkg1Handler1 {
	return &Pkg1Handler1Concrete{
		desc: "pkg1 handler1 instance",
	}
}
func (p *Pkg1Handler1Concrete) Get() interface{} {
	return p
}
func (p Pkg1Handler1Concrete) String() string {
	return p.desc
}
func (p *Pkg1Handler1Concrete) Foo1() string {
	return "foo1"
}

func GetPkg1Handler1FromCommon() (Pkg1Handler1, bool) {
	p, ok := common.GetHandler(Pkg1h1K)
	if !ok {
		return nil, ok
	}
	return p.(Pkg1Handler1), ok
}

func MustGetPkg1Handler1FromCommon() Pkg1Handler1 {
	return common.MustGetHandler(Pkg1h1K).(Pkg1Handler1)
}
