// +build:test
package common

func ResetHandlers() {
	hMap = &handlerMap{}
}

func CountInitHandlers() (count int) {
	initMap.Range(func(key, value any) bool {
		count += 1
		return true
	})
	return
}
