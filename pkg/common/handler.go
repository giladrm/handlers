package common

import (
	"sync"

	"github.com/pkg/errors"
)

type (
	// HandlerKey used to define actual handler key using embedding type of HandlerKey
	//
	// e.g.
	//	type myHandlerKey struct {
	//		commmon.HandlerKey
	//	}
	//
	// and instantiate a concrete key - must be exported in order
	// that other imports can use the handler after associating it with the key
	//
	//	var MyHandlerKey = MyHandlerKey{}
	HandlerKey interface {
		String() string
	}

	// RunHandler used to define actual handler using embedding type of RunHandler
	//
	// e.g.
	//	type MyHandlerInterface interface {
	//		commmon.RunHandler
	//	    ...
	//	}
	//
	// and then implement the interface (or skip the interface...)
	//
	//	type MyHandler struct {
	//		commmon.RunHandler
	//	    ...
	//	}
	RunHandler interface {
		Get() interface{}
		String() string
	}

	// InitHandler need to be implemented for the handler in order to have the handler initialized in a common manner
	//
	// e.g.
	//  type myInitHandler sturct{}
	//  func (i myInitHandler) Init(args ...interface{}) RunHandler {
	//      return &MyHander{}
	//   }
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
	ErrorKeyNotFound = errors.New("handler key not found")
	ErrorKeyExist    = errors.New("handler key already exist")
)

// handlerMap

func SetHandler(key HandlerKey, h RunHandler) {
	hMap.m.Store(key, h)
}

func (h *handlerMap) Range(f func(key HandlerKey, value RunHandler) bool) {
	h.m.Range(func(key, value any) bool {
		k := key.(HandlerKey)
		v := value.(RunHandler)
		f(k, v)
		return true
	})
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

func GetAllHandlers() (res map[HandlerKey]RunHandler) {
	res = map[HandlerKey]RunHandler{}
	hMap.Range(func(key HandlerKey, value RunHandler) bool { res[key] = value; return true })
	return
}

// initMap

// AddInitHandler register and associate a inithandler func with a handler key.
// for registration each handler within its file/package scope need to invokde this function
// in order to have seemless initialization
//
// e.g
//
//	func init() {
//	    common.AddInitHandler(GreenHouseFarmKey, greenHouseFarmInit{})
//	}
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
	hMap = &handlerMap{}
	initMap = &handlerInitMap{}
}

// InitAll initialize all registered handlers
func InitAll(args ...interface{}) {
	GetInitMap().Range(func(key, value any) bool {
		k := key.(HandlerKey)
		v := value.(InitHandler)
		SetHandler(k, v.Init(args))
		return true
	})
}

// InitSome initialize only request hanlders according to provided handlerKey list from the registered hadnlers
func InitSome(keys []HandlerKey, args ...interface{}) {
	kmap := map[HandlerKey]bool{}
	for _, e := range keys {
		kmap[e] = true
	}
	GetInitMap().Range(func(key, value any) bool {
		k := key.(HandlerKey)
		if _, ok := kmap[k]; ok {
			v := value.(InitHandler)
			SetHandler(k, v.Init(args))
		}
		return true
	})
}
