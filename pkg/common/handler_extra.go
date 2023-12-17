// +build:test
package common

func ResetHandlers() {
	hMap = &handlerMap{}
}

func CountInitHandlers() (count int) {
	initMap.traverse(func(_ HandlerKey, _ InitHandler) bool {
		count += 1
		return true
	})
	return
}
