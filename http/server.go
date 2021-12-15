package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
	Router(group string, rg ...gin.HandlerFunc) *gin.RouterGroup
	GetEngine() *gin.Engine
	GetRouterGroup(groupId string) *gin.RouterGroup
}

type Service struct {
	Name   string
	Port   int
	engine *gin.Engine
	groups map[string]*gin.RouterGroup
}

// NewHttp 创建http
func NewHttp(cc *Options) Server {
	app := &Service{}
	engine := gin.Default()
	app.groups = make(map[string]*gin.RouterGroup)
	app.engine = engine
	app.Name = cc.Service
	app.Port = cc.Port
	return app
}

// Router 创建router
func (app *Service) Router(groupId string, rg ...gin.HandlerFunc) *gin.RouterGroup {
	if app.engine == nil {
		panic("please provide a engine")
	}
	if app.groups == nil {
		panic("please provide a group")
	}
	group := app.engine.Group(groupId, rg...)
	app.groups[groupId] = group
	return group
}

// Run 运行服务
func (app *Service) Run() {
	if app.engine == nil {
		app.engine = gin.Default()
	}
	port := ":" + strconv.Itoa(app.Port)
	app.engine.Run(port)
}

func (app *Service) GetEngine() *gin.Engine {
	return app.engine
}

func (app *Service) GetRouterGroup(groupId string) *gin.RouterGroup {
	if _, ok := app.groups[groupId]; !ok {
		panic("group is not found")
		return nil
	}
	return app.groups[groupId]
}
