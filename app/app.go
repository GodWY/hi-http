package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/GodWY/hip/service"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Name   string
	Port   int
	engine *gin.Engine
	groups map[string]*gin.RouterGroup
	srv    *http.Server
}

// NewHttp 创建http
func NewHttp(cc *Options) service.Service {
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
	fmt.Println("[hip] running port :", port)
	app.engine.Run(port)

	// 使用http包监控端口.当失去连接时，平滑关闭
	srv := &http.Server{
		Addr:    ":8080",
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

// Close 关闭gin服务
func (app *Service) Close(cloSig chan bool) {

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if app.srv == nil {
		// TODO强制关不
		return
	}
	if err := app.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

// RegisterHttpHandler 注册http服务
func (app *Service) RegisterHttpHandler(group *gin.RouterGroup, topic string, handle *service.HandlerFunc) {
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

// RegisterPbHttpHandler 注册pb服务
func (app *Service) RegisterPbHttpHandler(group *gin.RouterGroup, topic string, handle *service.HandlerFunc) {
	return
}
