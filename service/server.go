package service

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc struct {
	Hf     gin.HandlerFunc
	Method string
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerHipFunc gin.HandlerFunc

type Service interface {
	RunHTTP()
	Router(group string, rg ...gin.HandlerFunc) *gin.RouterGroup
	GetRouterGroup(groupId string) *gin.RouterGroup
	RegisterHTTPHandler(group *gin.RouterGroup, topic string, handle *HandlerFunc)
	Close()
	RegisterPbHTTPHandler(group *gin.RouterGroup, topic string, handle *HandlerFunc)
	RegisterGroup(groupId string, handle ...HandlerHipFunc) *gin.RouterGroup
}

type Router interface {
	RegisterGetHandler(rounter string, handle ...HandlerHipFunc)
	RegisterPostHandler(rounter string, handle ...HandlerHipFunc)
	RegisterPutHandler(rounter string, handle ...HandlerHipFunc)
	RegisterAnyHandler(rounter string, handle ...HandlerHipFunc)
}
