package app1

import (
	"fmt"
	"handlers/pkg/store"
)

func App1() {
	appleStore := store.MustGetAppleStoreFromCommon()
	smith := appleStore.GetSmith()
	mangoStore := store.MustGetMangoStoreFromCommon()
	maya := mangoStore.GetMaya()
	fmt.Println(smith)
	fmt.Println(maya)
}
