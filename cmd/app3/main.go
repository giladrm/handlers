package main

import (
	"handlers/apps/app3"
	"handlers/pkg/common"
	"handlers/pkg/distribution"
)

func main() {
	requiredHandlerKeys := []common.HandlerKey{
		distribution.TruckDistributionKey,
	}
	common.InitSome(requiredHandlerKeys)
	app3.App3()
}
