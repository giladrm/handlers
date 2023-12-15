package app2

import (
	"fmt"
	"handlers/pkg/common"
	"handlers/pkg/distribution"
	"handlers/pkg/farm"
)

func App2() {
	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	orchad := farm.MustGetOrchadFarmFromCommon()
	truck := distribution.MustGetTruckDistributionFromCommon()

	fmt.Println(orchad.PickMango("maya"))
	fmt.Println(truck.Load("apple"))
}
