package main

import (
	"errors"
	"net/http"

	"github.com/GodWY/hip"
	"github.com/GodWY/hip/app"
	"github.com/GodWY/hip/proto/greeters"
	"github.com/gin-gonic/gin"
)

func main() {
	hip := hip.NewService(app.Builder().WithPort(7070).WithService("test"))
	greeters.RegisterGreeterHttpHandler(hip, &HelloWorld{})
	hip.Run()
	http.Get("")
}

type HelloWorld struct {
}

func (hw *HelloWorld) SayHello(ctx *gin.Context, in *greeters.SayHelloRequest) (out *greeters.SayHelloResponse, err error) {
	out = &greeters.SayHelloResponse{
		Msg: "success",
	}
	return
}
func (hw *HelloWorld) AskHello(ctx *gin.Context, in *greeters.AskHelloRequest) (out *greeters.AskHelloResponse, err error) {
	err = errors.New("sssss")
	return
}
