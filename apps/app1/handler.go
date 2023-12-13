package app1

import (
	"fmt"
	"handlers/pkg/common"
	_ "handlers/pkg/pkg1"
	_ "handlers/pkg/pkg2"
	_ "handlers/pkg/pkg3"
)

func init() {
	fmt.Println("app1 handler init")
	common.GetInitMap().Range(func(key, value any) bool {
		k := key.(common.HandlerKey)
		v := value.(common.InitHandler)
		fmt.Printf("key: %s, value: %+#v\n", k, v)
		common.SetHandler(k, v.Init())
		return true
	})
}
