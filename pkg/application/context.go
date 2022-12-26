package application

import (
	"errors"
	"gorm.io/gorm"
)

type Context struct {
	ctx map[string]*gorm.DB
}

func GetBean(name string) *gorm.DB {
	value, ok := ctx.ctx[name]

	if ok {
		return value
	} else {
		panic(errors.New("bean not found"))
		return value
	}

}

func SetBean(name string, bean *gorm.DB) {
	ctx.ctx[name] = bean
}

var ctx Context

func init() {
	ctx = Context{
		ctx: make(map[string]*gorm.DB),
	}
}
