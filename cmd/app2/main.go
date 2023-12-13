package main

import (
	"fmt"
	_ "handlers/apps/app2"
	"handlers/pkg/common"
	"handlers/pkg/pkg1"
	"handlers/pkg/pkg2"
)

func main() {
	common.InitAll()

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}
	p1 := pkg1.MustGetPkg1FromCommon()
	p2 := pkg2.MustGetPkg2FromCommon()
	fmt.Println(p1.Foo())
	fmt.Println(p2.Bar())

}
