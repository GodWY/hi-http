package main

import (
	"github.com/GodWY/hip"
	"github.com/GodWY/hip/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	p := hip.NewService(conf.WithPort(7070))
	e := gin.New()
	e.GET("/", func(context *gin.Context) {
		return
	})

	p.SetRouter(e)
	p.StartUp()
}


