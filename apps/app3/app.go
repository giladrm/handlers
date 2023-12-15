package app3

import (
	"fmt"
	"handlers/pkg/distribution"
)

func App3() {
	truck := distribution.MustGetTruckDistributionFromCommon()

	fmt.Println(truck.Load("something"))
}
