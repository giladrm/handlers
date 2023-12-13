package main

import (
	"fmt"
	_ "handlers/apps/app2"
	"handlers/pkg/common"
	"handlers/pkg/pkg1"
	"handlers/pkg/pkg2"
)

func main() {
	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}
	p1 := common.MustGetHandler(pkg1.Pkg1K).(pkg1.Pkg1)
	p2 := common.MustGetHandler(pkg2.Pkg2K).(pkg2.Pkg2)
	fmt.Println(p1.Foo())
	fmt.Println(p2.Bar())

}
