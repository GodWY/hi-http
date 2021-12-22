package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/GodWY/hip/service"
	"github.com/gin-gonic/gin"
)

// Service hip默认服务
type Service struct {
	Name   string // service name
	Port   int    // service port
	engine *gin.Engine
	groups map[string]*gin.RouterGroup
	srv    *http.Server
}

// NewHTTP 创建http
func NewHTTP(cc *Options) service.Service {
	app := new(Service)
	engine := gin.Default()
	app.groups = make(map[string]*gin.RouterGroup)
	app.engine = engine
	app.Name = cc.Service
	app.Port = cc.Port
	if !cc.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	return app
}

// Router 创建router
func (app *Service) Router(groupID string, rg ...gin.HandlerFunc) *gin.RouterGroup {
	if app.engine == nil {
		panic("please provide a engine")
	}
	if app.groups == nil {
		panic("please provide a group")
	}
	group := app.engine.Group(groupID, rg...)
	app.groups[groupID] = group
	return group
}

// RegisterGroup 注册服务组id
func (app *Service) RegisterGroup(groupID string, handle service.HandlerHipFunc) *gin.RouterGroup {
	if app.engine == nil {
		panic("please provide a engine")
	}
	if app.groups == nil {
		panic("please provide a group")
	}
	group := app.engine.Group(groupID, gin.HandlerFunc(handle))
	app.groups[groupID] = group
	return group
}

// Run 运行服务
func (app *Service) Run() {
	if app.engine == nil {
		panic("please create new Service")
	}
	port := ":" + strconv.Itoa(app.Port)
	app.engine.Run(port)
	// 使用http包监控端口.当失去连接时，平滑关闭
	srv := &http.Server{
		Addr:    port,
		Handler: app.engine,
	}
	app.srv = srv
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

// GetRouterGroup 通过组id获取实例
func (app *Service) GetRouterGroup(groupID string) *gin.RouterGroup {
	if _, ok := app.groups[groupID]; !ok {
		return nil
	}
	return app.groups[groupID]
}

// Close 关闭gin服务
func (app *Service) Close() {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if app.srv == nil {
		// TODO强制关不
		return
	}
	if err := app.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// RegisterHTTPHandler 注册http服务
func (app *Service) RegisterHTTPHandler(group *gin.RouterGroup, topic string, handle *service.HandlerFunc) {
	if handle == nil {
		panic("is error")
	}
	switch handle.Method {
	case "GET":
		group.GET(topic, handle.Hf)
	case "POST":
		group.POST(topic, handle.Hf)
	default:
		group.Any(topic, handle.Hf)
	}
}

// RegisterPbHTTPHandler 注册pb服务
func (app *Service) RegisterPbHTTPHandler(group *gin.RouterGroup, topic string, handle *service.HandlerFunc) {
	return
}
