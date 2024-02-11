package farm

import "handlers/pkg/common"

type (
	farmKey struct {
		keyString string
	}
	farmInit struct {
		initFunc func(args ...interface{}) common.RunHandler
	}
)

func (f farmKey) String() string                              { return f.keyString }
func (f farmInit) Init(args ...interface{}) common.RunHandler { return f.initFunc(args) }
