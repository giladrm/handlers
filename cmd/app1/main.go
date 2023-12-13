package main

import (
	"fmt"
	"handlers/pkg/common"
	"handlers/pkg/farm"
	"handlers/pkg/store"
)

func main() {
	requiredHandlerKeys := []common.HandlerKey{
		farm.OrchadFarmKey,
		store.AppleStoreKey,
		store.MangoStoreKey,
	}
	common.InitSome(requiredHandlerKeys)

	for k, v := range common.GetAllHandlers() {
		fmt.Printf("handler <%+#v>: <%+#v>\n", k, v.String())
	}
	appleStore := store.MustGetAppleStoreFromCommon()
	smith := appleStore.GetSmith()
	mangoStore := store.MustGetMangoStoreFromCommon()
	maya := mangoStore.GetMaya()
	fmt.Println(smith)
	fmt.Println(maya)

}
