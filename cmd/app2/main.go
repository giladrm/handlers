package main

import (
	"fmt"
	"handlers/apps/app2"
	"handlers/pkg/common"
)

func main() {
	common.InitAll()

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	app2.App2()
}
