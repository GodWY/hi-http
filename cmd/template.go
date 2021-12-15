package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemplateReq struct {
	H string
}

type TemplateRsp struct {
	Name string
}

type Template struct {
}

var T AckHttpService

type AckHttpService interface {
	HandleAck(ctx *gin.Context, in *TemplateReq) (out *TemplateRsp, err error)
	HandleAck1(ctx *gin.Context, in *TemplateReq) (out *TemplateRsp, err error)
}

type AckService interface {
	HandleAck(ctx *gin.Context)
	HandleAck1(ctx *gin.Context)
}

// RegisterHandler 注册服务
func registerHttpHandler(group *gin.RouterGroup, srv AckService) {
	group.Any("/ack", srv.HandleAck)
	group.Any("/ack1", srv.HandleAck1)
}

// RegisterPbHttpHandler 注册pb服务
func RegisterPbHttpHandler(group *gin.RouterGroup, srv AckHttpService) {
	tt := &Template{}
	registerHttpHandler(group, tt)
	T = srv
}

func (t *Template) HandleAck(ctx *gin.Context) {
	req := &TemplateReq{}
	if ok := ctx.Bind(req); ok != nil {
		ctx.JSON(http.StatusOK, "bind error")
		return
	}

	rsp, err := T.HandleAck(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (t *Template) HandleAck1(ctx *gin.Context) {
	req := &TemplateReq{}
	if ok := ctx.Bind(req); ok != nil {
		ctx.JSON(http.StatusOK, "bind error")
		return
	}

	rsp, err := T.HandleAck1(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}
