package main

import (
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

	app2.App2()
}
