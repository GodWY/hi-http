package app

import (
	"github.com/gin-gonic/gin"
)

type Options struct {
	Port       int
	Service    string
	MiddleWare []gin.HandlerFunc
	Debug      bool
}

func Builder() *Options {
	return &Options{}
}

// WithPort 生成Port
func (o *Options) WithPort(Port int) *Options {
	o.Port = Port
	return o
}

// WithService 服务的名字
func (o *Options) WithService(name string) *Options {
	o.Service = name
	return o
}

// WithMiddleWare
func (o *Options) WithMiddleWare(middleWare ...gin.HandlerFunc) *Options {
	newMiddWare := make([]gin.HandlerFunc, 0, len(middleWare))
	copy(newMiddWare, middleWare)
	o.MiddleWare = middleWare
	return o
}

// WithMiddleWare
func (o *Options) WithDebug(debug bool) *Options {
	o.Debug = debug
	return o
}
