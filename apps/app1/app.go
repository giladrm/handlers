package app1

import (
	"fmt"
	"handlers/pkg/distribution"
	"handlers/pkg/farm"
	"handlers/pkg/store"
)

func App1() {
	fruitFarm := farm.MustGetOrchadFarmFromCommon()
	deliveryTruck := distribution.MustGetTruckDistributionFromCommon()
	appleStore := store.MustGetAppleStoreFromCommon()
	mangoStore := store.MustGetMangoStoreFromCommon()

	pickedApples := fruitFarm.PickApple("smith")
	pickedMango := fruitFarm.PickMango("maya")
	loadingCargo := deliveryTruck.Load(pickedApples + " and " + pickedMango)
	unloadingCargo := deliveryTruck.Unload(pickedApples + " and " + pickedMango)
	smith := appleStore.GetSmith()
	maya := mangoStore.GetMaya()

	fmt.Println(pickedApples)
	fmt.Println(pickedMango)
	fmt.Println(loadingCargo)
	fmt.Println(unloadingCargo)
	fmt.Println(smith)
	fmt.Println(maya)
}
