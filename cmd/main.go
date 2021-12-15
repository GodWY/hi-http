package main

import (
	"github.com/GodWY/hi-http/internal/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cc := http.Builder().WithPort(7070)
	srv := http.NewHttp(cc)
	rg := srv.Router("api")
	RegisterPbHttpHandler(rg, &HelloWorld{})

	rg1 := srv.Router("api/v1")
	RegisterPbHttpHandler(rg1, &HelloWorld{})
	srv.Run()
}

type HelloWorld struct {
}

func (h *HelloWorld) HandleAck(ctx *gin.Context, req *TemplateReq) (rsp *TemplateRsp, err error) {
	rsp = &TemplateRsp{
		Name: "greet",
	}
	return
}

func (h *HelloWorld) HandleAck1(ctx *gin.Context, req *TemplateReq) (rsp *TemplateRsp, err error) {
	rsp = &TemplateRsp{
		Name: "greet1",
	}
	return
}
