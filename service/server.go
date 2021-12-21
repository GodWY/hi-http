package service

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc struct {
	Hf     gin.HandlerFunc
	Method string
}

type Service interface {
	Run()
	Router(group string, rg ...gin.HandlerFunc) *gin.RouterGroup
	GetEngine() *gin.Engine
	GetRouterGroup(groupId string) *gin.RouterGroup
	RegisterHttpHandler(group *gin.RouterGroup, topic string, handle *HandlerFunc)
	Close(cloSig chan bool)
	RegisterPbHttpHandler(group *gin.RouterGroup, topic string, handle *HandlerFunc)
}
