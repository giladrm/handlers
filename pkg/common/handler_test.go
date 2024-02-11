package common_test

import (
	"handlers/pkg/common"
	"testing"

	"github.com/pkg/errors"
	assert "github.com/stretchr/testify/require"
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
	testKeyType struct{}
	testInit    struct{}
)

func (k testKeyType) String() string { return "test key" }

func (i testInit) Init(args ...interface{}) common.RunHandler {
	return nil
}

func Test_handlers_adding_handler(t *testing.T) {
	defer func() {
		assert.NoError(t, common.RemoveInitHandler(testKeyType{}))
	}()
	testKey := testKeyType{}

	err := common.AddInitHandler(testKey, testInit{}, 1)
	assert.NoError(t, err)
	err = common.AddInitHandler(testKey, testInit{}, 1)
	assert.Error(t, err)
	assert.ErrorContains(t, err, common.ErrorKeyExist.Error())
	assert.ErrorContains(t, err, testKey.String())
}

func Test_handlers_one_not_exist(t *testing.T) {
	testKey := testKeyType{}

	assert.PanicsWithError(t, errors.Wrap(common.ErrorKeyNotFound, testKey.String()).Error(), func() {
		common.InitSome([]common.HandlerKey{testKey})
	})
}

func Test_handlers_one_exist(t *testing.T) {
	defer TearDown()
	assert.NotPanics(t, func() {
		common.InitSome([]common.HandlerKey{FooKey})
	})

	res := common.GetAllHandlers()
	assert.Len(t, res, 1)

	h, ok := common.GetHandler(FooKey)
	assert.True(t, ok)
	assert.NotNil(t, h)
	assert.Contains(t, h.String(), FooKey.String())

	h2, ok := common.GetHandler(BarKey)
	assert.False(t, ok)
	assert.Nil(t, h2)
}

func Test_handlers_Get_invalid_handler_type(t *testing.T) {
	defer func() {
		assert.NoError(t, common.RemoveInitHandler(FooKey))
		assert.NoError(t, common.AddInitHandler(FooKey, fooInit{}, 1))
	}()
	assert.NoError(t, common.RemoveInitHandler(FooKey))
	err := common.AddInitHandler(FooKey, fizzInit{}, 1)
	assert.NoError(t, err)

	defer TearDown()
	handlerKeyList := []common.HandlerKey{
		FooKey,
	}
	common.InitSome(handlerKeyList)

	h, ok := GetTestFooFromCommon()
	assert.False(t, ok)
	assert.Nil(t, h)
}

func Test_handlers_MustGet_init_some(t *testing.T) {
	defer TearDown()
	handlerKeyList := []common.HandlerKey{
		FooKey,
		BarKey,
		FooBarKey,
	}
	common.InitSome(handlerKeyList)

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

func Test_handlers_MustGet_init_all(t *testing.T) {
	defer TearDown()
	common.InitAll()

	res := common.GetAllHandlers()
	assert.Len(t, res, common.CountInitHandlers())

	var fizzbuzz TestFizzBuzz
	assert.Nil(t, fizzbuzz)
	assert.NotPanics(t, func() { fizzbuzz = MustGetTestFizzBuzzFromCommon() })
	assert.NotNil(t, fizzbuzz)

	var buzz TestBuzz
	assert.Nil(t, buzz)
	assert.NotPanics(t, func() { buzz = MustGetTestBuzzFromCommon() })
	assert.NotNil(t, buzz)
}
