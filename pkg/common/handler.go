package common

import (
	"sort"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

// HandlerKey used to define actual handler key using embedding type of HandlerKey
//
// e.g.
//
//	type myHandlerKey struct {
//		commmon.HandlerKey
//	}
//
// and instantiate a concrete key - must be exported in order
// that other imports can use the handler after associating it with the key
//
//	var MyHandlerKey = myHandlerKey{}
type HandlerKey interface {
	String() string
}

// RunHandler used to define actual handler using embedding type of RunHandler
//
// e.g.
//
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
type RunHandler interface {
	Get() interface{} //optional
	String() string
}

// InitHandler need to be implemented for the handler in order to have the handler initialized in a common manner
//
// e.g.
//
//	type myInitHandler sturct{}
//	func (i myInitHandler) Init(args ...interface{}) RunHandler {
//	    return &MyHander{}
//	 }
type InitHandler interface {
	Init(args ...interface{}) RunHandler
}

// internal managing
type (
	handlerInitDuo struct {
		key     HandlerKey
		handler InitHandler
	}
	handlerMap struct {
		m sync.Map
	}
	handlerInitMap struct {
		m map[int][]handlerInitDuo
	}
)

var (
	hMap             *handlerMap
	initMap          *handlerInitMap
	ErrorKeyNotFound = errors.New("handler key not found")
	ErrorKeyExist    = errors.New("handler key already exist")
)

// handlerMap

func SetHandler(key HandlerKey, h RunHandler)                     { hMap.m.Store(key, h) }
func RangeHandlers(f func(key HandlerKey, value RunHandler) bool) { hMap.traverse(f) }
func GetHandler(key HandlerKey) (RunHandler, bool)                { return hMap.getHandler(key) }
func MustGetHandler(key HandlerKey) RunHandler                    { return hMap.mustGetHandler(key) }
func GetAllHandlers() map[HandlerKey]RunHandler                   { return hMap.getAllHandlers() }

func (h *handlerMap) traverse(f func(key HandlerKey, value RunHandler) bool) {
	h.m.Range(func(key, value any) bool {
		k := key.(HandlerKey)
		v := value.(RunHandler)
		f(k, v)
		return true
	})
}

func (h *handlerMap) getHandler(key HandlerKey) (RunHandler, bool) {
	handler, ok := h.m.Load(key)
	if ok {
		return handler.(RunHandler), ok
	}
	return nil, ok
}

func (h *handlerMap) mustGetHandler(key HandlerKey) RunHandler {
	handler, ok := h.getHandler(key)
	if !ok {
		panic(errors.Wrap(ErrorKeyNotFound, key.String()))
	}
	return handler
}

func (h *handlerMap) getAllHandlers() (res map[HandlerKey]RunHandler) {
	res = map[HandlerKey]RunHandler{}
	hMap.traverse(func(key HandlerKey, value RunHandler) bool { res[key] = value; return true })
	return
}

// initMap
func (h *handlerInitMap) existInitHandler(key HandlerKey) error {
	for _, v := range h.m {
		for _, e := range v {
			if e.key == key {
				return errors.Wrap(ErrorKeyExist, key.String())
			}
		}
	}
	return nil
}

func (h *handlerInitMap) addInitHandler(key HandlerKey, initHandler InitHandler, phaseIndex int) error {
	if err := h.existInitHandler(key); err != nil {
		return err
	}
	duo := handlerInitDuo{key: key, handler: initHandler}
	list, ok := h.m[phaseIndex]
	if ok {
		list = append(list, duo)
	} else {
		list = []handlerInitDuo{duo}
	}
	h.m[phaseIndex] = list
	return nil
}

func (h *handlerInitMap) traverseByPhase(f func(key HandlerKey, value InitHandler) bool) {
	keys := maps.Keys(h.m)
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, k := range keys {
		for _, e := range h.m[k] {
			f(e.key, e.handler)
		}
	}
}

func (h *handlerInitMap) initAll(args ...interface{}) {
	h.traverseByPhase(func(key HandlerKey, value InitHandler) bool {
		SetHandler(key, value.Init(args))
		return true
	})
}

func (h *handlerInitMap) initSome(keys []HandlerKey, args ...interface{}) {
	kmap := map[HandlerKey]bool{}
	for _, e := range keys {
		kmap[e] = true
	}
	h.traverseByPhase(func(key HandlerKey, value InitHandler) bool {
		if _, ok := kmap[key]; ok {
			SetHandler(key, value.Init(args))
			kmap[key] = false
		}
		return true
	})
	for k, v := range kmap {
		if v {
			panic(errors.Wrap(ErrorKeyNotFound, k.String()))
		}
	}
}

// AddInitHandler register and associate a inithandler func with a handler key.
//
// for registration each handler within its file/package scope need to invokde this function
// in order to have seemless initialization
//
// e.g
//
//	func init() {
//	    common.AddInitHandler(GreenHouseFarmKey, greenHouseFarmInit{})
//	}
func AddInitHandler(key HandlerKey, initHandler InitHandler, phaseIndex int) error {
	return initMap.addInitHandler(key, initHandler, phaseIndex)
}

// InitAll initialize all registered handlers
//
// usefull for most cases where there is no interdependency between handlers
func InitAll(args ...interface{}) { initMap.initAll(args) }

// InitSome initialize only requested handlers according to provided handlerKey list from the registered hadnlers
//
// usefull for cases where there is an interdependency berween handlers
func InitSome(keys []HandlerKey, args ...interface{}) { initMap.initSome(keys, args) }

func init() {
	hMap = &handlerMap{}
	initMap = &handlerInitMap{map[int][]handlerInitDuo{}}
}
