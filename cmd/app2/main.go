package main

import (
	"fmt"
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

	orchad := farm.MustGetOrchadFarmFromCommon()
	truck := distribution.MustGetTruckDistributionFromCommon()

	fmt.Println(orchad.PickApple("maya"))
	fmt.Println(truck.Load("apple"))

}
