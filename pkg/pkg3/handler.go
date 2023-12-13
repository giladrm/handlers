package pkg3

import "handlers/pkg/common"

type (
	Pkg3 interface {
		common.RunHandler
		Baz() string
	}

	Pkg3Concrete struct {
		common.RunHandler
		desc string
	}
)

func NewPkg3(args ...interface{}) Pkg3 {
	return &Pkg3Concrete{
		desc: "pkg3 instance",
	}
}
func (p *Pkg3Concrete) Get() interface{} {
	return p
}
func (p Pkg3Concrete) String() string {
	return p.desc
}
func (p *Pkg3Concrete) Baz() string {
	return "baza"
}

func MustGetPkg3FromCommon() Pkg3 {
	return common.MustGetHandler(Pkg3K).(Pkg3)
}
