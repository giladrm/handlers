package main

import (
	"fmt"
	"handlers/apps/app2"
	"handlers/pkg/common"
	"handlers/pkg/distribution"
	"handlers/pkg/farm"
)

func main() {
	requiredHandlerKeys := []common.HandlerKey{
		farm.OrchadFarmKey,
		distribution.TruckDistributionKey,
	}
	common.InitSome(requiredHandlerKeys)

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	app2.App2()
}
