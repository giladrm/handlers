package main

import (
	"fmt"
	"handlers/apps/app3"
	"handlers/pkg/common"
)

func main() {
	common.InitAll()

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	app3.App3()
}
