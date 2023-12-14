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

var (
	hMap             *handlerMap
	initMap          *handlerInitMap
	ErrorKeyNotFound = errors.New("key not found")
	ErrorKeyExist    = errors.New("key already exist")
)

// handlerMap

func SetHandler(key HandlerKey, h RunHandler) {
	hMap.m.Store(key, h)
}

func (h *handlerMap) Range(f func(key, value any) bool) {
	h.m.Range(f)
}

func GetHandlersMap() *handlerMap {
	return hMap
}

func GetHandler(key HandlerKey) (RunHandler, bool) {
	h, ok := hMap.m.Load(key)
	if ok {
		return h.(RunHandler), ok
	}
	return nil, ok
}

func MustGetHandler(key HandlerKey) RunHandler {
	h, ok := hMap.m.Load(key)
	if !ok {
		panic(errors.Wrap(ErrorKeyNotFound, key.String()))
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

func AddInitHandler(key HandlerKey, initFunc InitHandler) (err error) {
	for _, e := range initMap.m {
		if e.k == key {
			return errors.Wrap(err, key.String())
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
	InitHandlers()
}

func InitHandlers() {
	hMap = &handlerMap{}
	initMap = &handlerInitMap{}
}

func InitAll(args ...interface{}) {
	GetInitMap().Range(func(key, value any) bool {
		k := key.(HandlerKey)
		v := value.(InitHandler)
		// fmt.Printf("key: %s, value: %+#v\n", k, v)
		SetHandler(k, v.Init(args))
		return true
	})
}

func InitSome(keys []HandlerKey, args ...interface{}) {
	kmap := map[HandlerKey]bool{}
	for _, e := range keys {
		kmap[e] = true
	}
	GetInitMap().Range(func(key, value any) bool {
		k := key.(HandlerKey)
		if _, ok := kmap[k]; ok {
			v := value.(InitHandler)
			// fmt.Printf("key: %s, value: %+#v\n", k, v)
			SetHandler(k, v.Init(args))
		}
		return true
	})
}
