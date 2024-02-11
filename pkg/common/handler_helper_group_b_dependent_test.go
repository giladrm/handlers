package common_test

import (
	"handlers/pkg/common"
)

// foobar object

type (
	TestFooBar interface {
		common.RunHandler
		FooBar(string) string
	}

	TestFooBarHandler struct {
		common.RunHandler
		desc string
		foo  TestFoo
		bar  TestBar
	}

	fooBarKey  struct{ common.HandlerKey }
	fooBarInit struct{}
)

func (k fooBarKey) String() string { return "foobar key" }

var (
	FooBarKey = fooBarKey{}
)

func NewFooBar(args ...interface{}) TestFooBar {
	return &TestFooBarHandler{
		desc: "foobar test instance",
		foo:  MustGetTestFooFromCommon(),
		bar:  MustGetTestBarFromCommon(),
	}
}

func (p *TestFooBarHandler) String() string   { return p.desc }
func (p *TestFooBarHandler) Get() interface{} { return p }
func (p *TestFooBarHandler) FooBar(inTest string) string {
	return p.foo.Foo(inTest) + " " + p.bar.Bar(inTest)
}

func GetTestFooBarFromCommon() (TestFooBar, bool) {
	p, ok := common.GetHandler(FooBarKey)
	if !ok {
		return nil, ok
	}
	return p.(TestFooBar), ok
}

func MustGetTestFooBarFromCommon() TestFooBar {
	return common.MustGetHandler(FooBarKey).(TestFooBar)
}

func (a fooBarInit) Init(args ...interface{}) common.RunHandler {
	return NewFooBar(args)
}

func init() {
	common.AddInitHandler(FooBarKey, fooBarInit{}, 2)
}

// fizzbuzz object

type (
	TestFizzBuzz interface {
		common.RunHandler
		FizzBuzz(string) string
	}

	TestFizzBuzzHandler struct {
		common.RunHandler
		desc string
		fizz TestFizz
		buzz TestBuzz
	}

	fizzBuzzKey  struct{ common.HandlerKey }
	fizzBuzzInit struct{}
)

func (k fizzBuzzKey) String() string { return "fizzbuzz key" }

var (
	FizzBuzzKey = fizzBuzzKey{}
)

func NewFizzBuzz(args ...interface{}) TestFizzBuzz {
	return &TestFizzBuzzHandler{
		desc: "fizzbuzz test instance",
		fizz: MustGetTestFizzFromCommon(),
		buzz: MustGetTestBuzzFromCommon(),
	}
}

func (p *TestFizzBuzzHandler) String() string   { return p.desc }
func (p *TestFizzBuzzHandler) Get() interface{} { return p }
func (p *TestFizzBuzzHandler) FizzBuzz(inTest string) string {
	return p.fizz.Fizz(inTest) + " " + p.buzz.Buzz(inTest)
}

func GetTestFizzBuzzFromCommon() (TestFizzBuzz, bool) {
	p, ok := common.GetHandler(FizzBuzzKey)
	if !ok {
		return nil, ok
	}
	return p.(TestFizzBuzz), ok
}

func MustGetTestFizzBuzzFromCommon() TestFizzBuzz {
	return common.MustGetHandler(FizzBuzzKey).(TestFizzBuzz)
}

func (a fizzBuzzInit) Init(args ...interface{}) common.RunHandler {
	return NewFizzBuzz(args)
}

func init() {
	common.AddInitHandler(FizzBuzzKey, fizzBuzzInit{}, 2)
}
