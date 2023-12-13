package common

import (
	"sync"

	"github.com/pkg/errors"
)

type (
	HandlerKey interface {
		String() string
	}
	RunHandler interface {
		Get() interface{}
		String() string
	}

	InitHandler interface {
		Init(args ...interface{}) RunHandler
	}
	handlerMap struct {
		m sync.Map
	}
	handlerInitMap struct {
		m []struct {
			k HandlerKey
			v InitHandler
		}
	}
)

// handlerMap
var hMap *handlerMap

func SetHandler(key HandlerKey, h RunHandler) {
	hMap.m.Store(key, h)
}
func GetHandler(key HandlerKey) (RunHandler, bool) {
	h, ok := hMap.m.Load(key)
	return h.(RunHandler), ok
}
func MustGetHandler(key HandlerKey) RunHandler {
	h, ok := hMap.m.Load(key)
	if !ok {
		panic("key not found")
	}
	return h.(RunHandler)
}

func GetAllHandlers() (res map[string]RunHandler) {
	res = map[string]RunHandler{}
	hMap.m.Range(func(key, value any) bool {
		k := key.(HandlerKey)
		res[k.String()] = value.(RunHandler)
		return true
	})
	return
}

// initMap
var initMap *handlerInitMap

func AddInitHandler(key HandlerKey, initFunc InitHandler) (err error) {
	for _, e := range initMap.m {
		if e.k == key {
			return errors.Errorf("key %v, already exists", key)
		}
	}
	initMap.m = append(initMap.m, struct {
		k HandlerKey
		v InitHandler
	}{key, initFunc})
	return
}

func GetInitMap() *handlerInitMap {
	return initMap
}

func (h *handlerInitMap) Range(f func(key any, value any) bool) {
	for _, e := range h.m {
		f(e.k, e.v)
	}
}

func init() {
	hMap = &handlerMap{}
	initMap = &handlerInitMap{}
}
func InitAll() {
	GetInitMap().Range(func(key, value any) bool {
		k := key.(HandlerKey)
		v := value.(InitHandler)
		// fmt.Printf("key: %s, value: %+#v\n", k, v)
		SetHandler(k, v.Init())
		return true
	})
}
