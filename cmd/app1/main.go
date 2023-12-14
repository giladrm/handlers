package main

import (
	"fmt"
	"handlers/apps/app1"
	"handlers/pkg/common"
	"handlers/pkg/distribution"
	"handlers/pkg/farm"
	"handlers/pkg/store"
)

func main() {
	requiredHandlerKeys := []common.HandlerKey{
		farm.OrchadFarmKey,
		distribution.TruckDistributionKey,
		store.AppleStoreKey,
		store.MangoStoreKey,
	}
	common.InitSome(requiredHandlerKeys)

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}
	app1.App1()
}
