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
	p1 := common.MustGetHandler(pkg1.Pkg1K).(pkg1.Pkg1)
	p3 := common.MustGetHandler(pkg3.Pkg3K).(pkg3.Pkg3)
	fmt.Println(p1.Foo())
	fmt.Println(p3.Baz())

}
