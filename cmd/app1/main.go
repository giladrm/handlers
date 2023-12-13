package main

import (
	"fmt"
	_ "handlers/apps/app1"
	"handlers/pkg/common"
	"handlers/pkg/pkg1"
	"handlers/pkg/pkg2"
	"handlers/pkg/pkg3"
)

func main() {
	common.InitAll()

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}
	p1h1 := pkg1.MustGetPkg1Handler1FromCommon()
	p1h2 := pkg1.MustGetPkg1Handler2FromCommon()
	p2 := pkg2.MustGetPkg2FromCommon()
	p3 := pkg3.MustGetPkg3FromCommon()
	fmt.Println(p1h1.Foo1())
	fmt.Println(p1h2.FooBar())
	fmt.Println(p2.Bar())
	fmt.Println(p3.Baz())

}
