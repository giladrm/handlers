package common_test

import (
	"handlers/pkg/common"
)

// foo object

type (
	TestFoo interface {
		common.RunHandler
		Foo(string) string
	}

	TestFooHandler struct {
		common.RunHandler
		desc string
	}

	fooKey  struct{ common.HandlerKey }
	fooInit struct{}
)

func (k fooKey) String() string { return "foo key" }

var (
	FooKey = fooKey{}
)

func NewFoo(args ...interface{}) TestFoo {
	return &TestFooHandler{
		desc: FooKey.String() + " test instance",
	}
}

func (p *TestFooHandler) String() string           { return p.desc }
func (p *TestFooHandler) Get() interface{}         { return p }
func (p *TestFooHandler) Foo(inTest string) string { return "foo " + inTest }

func GetTestFooFromCommon() (TestFoo, bool) {
	p, ok := common.GetHandler(FooKey)
	if !ok {
		return nil, ok
	}
	return p.(TestFoo), ok
}

func MustGetTestFooFromCommon() TestFoo {
	return common.MustGetHandler(FooKey).(TestFoo)
}

func (a fooInit) Init(args ...interface{}) common.RunHandler {
	return NewFoo(args)
}

func init() {
	common.AddInitHandler(FooKey, fooInit{}, 1)
}

// bar object

type (
	TestBar interface {
		common.RunHandler
		Bar(string) string
	}

	TestBarHandler struct {
		common.RunHandler
		desc string
	}

	barKey  struct{ common.HandlerKey }
	barInit struct{}
)

func (k barKey) String() string { return "bar key" }

var (
	BarKey = barKey{}
)

func NewBar(args ...interface{}) TestBar {
	return &TestBarHandler{
		desc: "bar test instance",
	}
}

func (p *TestBarHandler) String() string           { return p.desc }
func (p *TestBarHandler) Get() interface{}         { return p }
func (p *TestBarHandler) Bar(inTest string) string { return "bar " + inTest }

func GetTestBarFromCommon() (TestBar, bool) {
	p, ok := common.GetHandler(BarKey)
	if !ok {
		return nil, ok
	}
	return p.(TestBar), ok
}

func MustGetTestBarFromCommon() TestBar {
	return common.MustGetHandler(BarKey).(TestBar)
}

func (a barInit) Init(args ...interface{}) common.RunHandler {
	return NewBar(args)
}

func init() {
	common.AddInitHandler(BarKey, barInit{}, 1)
}

// fizz object

type (
	TestFizz interface {
		common.RunHandler
		Fizz(string) string
	}

	TestFizzHandler struct {
		common.RunHandler
		desc string
	}

	fizzKey  struct{ common.HandlerKey }
	fizzInit struct{}
)

func (k fizzKey) String() string { return "fizz key" }

var (
	FizzKey = fizzKey{}
)

func NewFizz(args ...interface{}) TestFizz {
	return &TestFizzHandler{
		desc: "fizz test instance",
	}
}

func (p *TestFizzHandler) String() string            { return p.desc }
func (p *TestFizzHandler) Get() interface{}          { return p }
func (p *TestFizzHandler) Fizz(inTest string) string { return "fizz " + inTest }

func GetTestFizzFromCommon() (TestFizz, bool) {
	p, ok := common.GetHandler(FizzKey)
	if !ok {
		return nil, ok
	}
	return p.(TestFizz), ok
}

func MustGetTestFizzFromCommon() TestFizz {
	return common.MustGetHandler(FizzKey).(TestFizz)
}

func (a fizzInit) Init(args ...interface{}) common.RunHandler {
	return NewFizz(args)
}

func init() {
	common.AddInitHandler(FizzKey, fizzInit{}, 1)
}

// buzz object

type (
	TestBuzz interface {
		common.RunHandler
		Buzz(string) string
	}

	TestBuzzHandler struct {
		common.RunHandler
		desc string
	}

	buzzKey  struct{ common.HandlerKey }
	buzzInit struct{}
)

func (k buzzKey) String() string { return "buzz key" }

var (
	BuzzKey = buzzKey{}
)

func NewBuzz(args ...interface{}) TestBuzz {
	return &TestBuzzHandler{
		desc: "buzz test instance",
	}
}

func (p *TestBuzzHandler) String() string            { return p.desc }
func (p *TestBuzzHandler) Get() interface{}          { return p }
func (p *TestBuzzHandler) Buzz(inTest string) string { return "buzz " + inTest }

func GetTestBuzzFromCommon() (TestBuzz, bool) {
	p, ok := common.GetHandler(BuzzKey)
	if !ok {
		return nil, ok
	}
	return p.(TestBuzz), ok
}

func MustGetTestBuzzFromCommon() TestBuzz {
	return common.MustGetHandler(BuzzKey).(TestBuzz)
}

func (a buzzInit) Init(args ...interface{}) common.RunHandler {
	return NewBuzz(args)
}

func init() {
	common.AddInitHandler(BuzzKey, buzzInit{}, 1)
}
