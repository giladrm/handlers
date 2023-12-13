package pkg1

import "handlers/pkg/common"

type (
	Pkg1Handler2 interface {
		common.RunHandler
		FooBar() string
	}

	Pkg1Handler2Concrete struct {
		common.RunHandler
		desc string
	}
)

func NewPkg1Handler2(args ...interface{}) Pkg1Handler2 {
	return &Pkg1Handler2Concrete{
		desc: "pkg1handler2 instance",
	}
}
func (p *Pkg1Handler2Concrete) Get() interface{} {
	return p
}
func (p Pkg1Handler2Concrete) String() string {
	return p.desc
}
func (p *Pkg1Handler2Concrete) FooBar() string {
	return "foobar"
}

func GetPkg1Handler2FromCommon() (Pkg1Handler2, bool) {
	p, ok := common.GetHandler(Pkg1h2K)
	if !ok {
		return nil, ok
	}
	return p.(Pkg1Handler2), ok
}

func MustGetPkg1Handler2FromCommon() Pkg1Handler2 {
	return common.MustGetHandler(Pkg1h2K).(Pkg1Handler2)
}
