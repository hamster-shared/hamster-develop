package application

import "errors"

type Context[T any] struct {
	ctx map[string]T
}

func GetBean[T any](name string) T {
	value, ok := ctx.ctx[name]

	if ok {
		return value.(T)
	} else {
		panic(errors.New("bean not found"))
		return value.(T)
	}

}

func SetBean[T any](name string, bean T) {
	ctx.ctx[name] = bean
}

var ctx Context[any]

func init() {
	ctx = Context[any]{
		ctx: make(map[string]any),
	}
}
