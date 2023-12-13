package main

import (
	"fmt"
	_ "handlers/apps/app3"
	"handlers/pkg/common"
	"handlers/pkg/pkg1"
	"handlers/pkg/pkg3"
)

func main() {
	common.InitAll()

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	p1 := pkg1.MustGetPkg1FromCommon()
	p3 := pkg3.MustGetPkg3FromCommon()
	fmt.Println(p1.Foo())
	fmt.Println(p3.Baz())

}
