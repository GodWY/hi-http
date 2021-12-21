package main

import (
	"github.com/GodWY/hip"
	"github.com/GodWY/hip/app"
	"github.com/GodWY/hip/greeter"
	"github.com/gin-gonic/gin"
)

func main() {
	hip := hip.NewService(app.Builder().WithPort(7070).WithService("test"))
	greeter.RegisterGreeterHttpHandler(hip, &HelloWorld{})
	hip.Run()
}

type HelloWorld struct {
}

func (hl *HelloWorld) Hello(ctx *gin.Context, in *greeter.Request) (out *greeter.Response, err error) {
	out = &greeter.Response{
		Msg: "success",
	}
	return
}
func (hl *HelloWorld) Stream(ctx *gin.Context, in *greeter.Request) (out *greeter.Response, err error) {
	out = &greeter.Response{
		Msg: "success",
	}
	return
}
