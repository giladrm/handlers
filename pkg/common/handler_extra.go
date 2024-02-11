// +build:test
package common

func ResetHandlers() {
	hMap = &handlerMap{}
}

func CountInitHandlers() (count int) {
	initMap.traverseByPhase(func(_ HandlerKey, _ InitHandler) bool {
		count += 1
		return true
	})
	return
}

func (h *handlerInitMap) removeInitHandler(key HandlerKey) error {
	for k, v := range h.m {
		for i, e := range v {
			if e.key == key {
				h.m[k] = append(h.m[k][:i], h.m[k][i+1:]...)
				return nil
			}
		}
	}
	return ErrorKeyExist
}

func RemoveInitHandler(key HandlerKey) error {
	return initMap.removeInitHandler(key)
}
