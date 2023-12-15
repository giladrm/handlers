package app3

import (
	"fmt"
	"handlers/pkg/common"
	"handlers/pkg/distribution"
)

func App3() {
	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}

	truck := distribution.MustGetTruckDistributionFromCommon()

	fmt.Println(truck.Load("something"))
}
