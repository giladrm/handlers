package main

import (
	"fmt"
	"handlers/apps/app3"
	"handlers/pkg/common"
	"handlers/pkg/distribution"
)

func main() {
	requiredHandlerKeys := []common.HandlerKey{
		distribution.TruckDistributionKey,
	}
	common.InitSome(requiredHandlerKeys)

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	app3.App3()
}
