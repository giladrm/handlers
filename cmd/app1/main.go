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
	p1 := common.MustGetHandler(pkg1.Pkg1K).(pkg1.Pkg1)
	p2 := common.MustGetHandler(pkg2.Pkg2K).(pkg2.Pkg2)
	p3 := common.MustGetHandler(pkg3.Pkg3K).(pkg3.Pkg3)
	fmt.Println(p1.Foo())
	fmt.Println(p2.Bar())
	fmt.Println(p3.Baz())

}
