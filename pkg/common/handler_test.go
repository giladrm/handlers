package common_test

import (
	"handlers/pkg/common"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TearDown() {
	common.ResetHandlers()
}

func Test_handlers(t *testing.T) {
	common.InitAll()
	defer TearDown()

	res := common.GetAllHandlers()
	assert.Len(t, res, common.CountInitHandlers())
	common.RangeHandlers(func(key common.HandlerKey, value common.RunHandler) bool {
		t.Log(value.String())
		return true
	})
}

type (
	testKeyType struct{ common.HandlerKey }
	testInit    struct{}
)

func (i testInit) Init(args ...interface{}) common.RunHandler {
	return nil
}

func Test_handlers_adding_handler(t *testing.T) {
	defer TearDown()
	testKey := testKeyType{}

	assert.NotPanicsf(t, func() {
		common.AddInitHandler(testKey, testInit{})
	}, "should not panic on addInitHandler for first time use key")
	assert.Panicsf(t, func() {
		common.AddInitHandler(testKey, testInit{})
	}, "should panic on addInitHandler for consecutive use of already used key")

}

func Test_handlers_one(t *testing.T) {
	common.InitSome([]common.HandlerKey{FooKey})
	defer TearDown()

	res := common.GetAllHandlers()
	assert.Len(t, res, 1)

	h, ok := common.GetHandler(FooKey)
	assert.True(t, ok)
	assert.NotNil(t, h)
	assert.Contains(t, h.String(), "foo")

	h2, ok := common.GetHandler(BarKey)
	assert.False(t, ok)
	assert.Nil(t, h2)
}

func Test_handlers_MustGet(t *testing.T) {
	handlerKeyList := []common.HandlerKey{
		FooKey,
		BarKey,
		FooBarKey,
	}
	common.InitSome(handlerKeyList)
	defer TearDown()

	res := common.GetAllHandlers()
	assert.Len(t, res, len(handlerKeyList))

	assert.PanicsWithError(t,
		errors.Wrap(common.ErrorKeyNotFound, FizzKey.String()).Error(),
		func() { common.MustGetHandler(FizzKey) },
	)

	var foobar TestFooBar
	assert.Nil(t, foobar)
	assert.NotPanics(t, func() { foobar = MustGetTestFooBarFromCommon() })
	assert.NotNil(t, foobar)

	var bar TestBar
	assert.Nil(t, bar)
	assert.NotPanics(t, func() { bar = MustGetTestBarFromCommon() })
	assert.NotNil(t, bar)
}
