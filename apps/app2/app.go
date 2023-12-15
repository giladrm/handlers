package app2

import (
	"fmt"
	"handlers/pkg/distribution"
	"handlers/pkg/farm"
)

func App2() {
	orchad := farm.MustGetOrchadFarmFromCommon()
	truck := distribution.MustGetTruckDistributionFromCommon()

	fmt.Println(orchad.PickMango("maya"))
	fmt.Println(truck.Load("apple"))
}
