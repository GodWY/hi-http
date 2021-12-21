package template

var html = `
package {{.package}}

import (
	"net/http"

	"github.com/GodWY/hip/service"
	"github.com/gin-gonic/gin"
)

var ts {{.service}}


type {{.service}} interface {
{{- range $key, $value := .Topic }}
	{{$value}}
{{- end}}
}

type AckService interface {
{{- range $key, $value := .Topic }}
	{{$value}}
{{- end}}
}

// RegisterHandler 注册服务
func registerHttpHandler(srv service.Service, srvs AckService) {
	// 注册服务组
	group := srv.Router("ack")
	{{- range $key, $value := .Topic }}
	group.Any($key,$value)
	{{- end}}
}

type hip struct{}

// RegisterPbHttpHandler 注册pb服务
func RegisterPbHttpHandler(srv service.Service, srvs {{.service}}) {
	hi := new(hip)
	ts = srvs
	registerHttpHandler(srv, ts)
}


{{- range $key, $value := .Topic }}
func (h *hip) Handle{{.h}}(ctx *gin.Context) {
	req := &TemplateReq{}
	if ok := ctx.Bind(req); ok != nil {
		ctx.JSON(http.StatusOK, "bind error")
		return
	}
	rsp, err := ts.HandleAck(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}
{{- end}}
`
