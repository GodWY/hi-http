package service

import (
	"net/http"
)

type ServerInfo struct {
	// ServiceName 服务名称
	ServiceName string
}
type Service interface {
	OnInit()
	StartUp() error
	Stop()
	SetRouter(h http.Handler)
}
