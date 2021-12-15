package main

import (
	"github.com/GodWY/hi-http/http"
	"github.com/gin-gonic/gin"
)

func main() {
	cc := http.Builder().WithPort(7070)
	srv := http.NewHttp(cc)
	rg := srv.Router("api")
	http.RegisterPbHttpHandler(rg, &HelloWorld{})

	rg1 := srv.Router("api/v1")
	http.RegisterPbHttpHandler(rg1, &HelloWorld{})
	srv.Run()
}

type HelloWorld struct {
}

func (h *HelloWorld) HandleAck(ctx *gin.Context, req *http.TemplateReq) (rsp *http.TemplateRsp, err error) {
	rsp = &http.TemplateRsp{
		Name: "greet",
	}
	return
}

func (h *HelloWorld) HandleAck1(ctx *gin.Context, req *http.TemplateReq) (rsp *http.TemplateRsp, err error) {
	rsp = &http.TemplateRsp{
		Name: "greet1",
	}
	return
}
