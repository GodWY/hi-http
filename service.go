package hip

import (
	"context"
	"fmt"
	"github.com/GodWY/hip/conf"
	"github.com/GodWY/hip/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"syscall"
)

// Service hip默认服务
type server struct {
	ctx  context.Context
	srv  *http.Server
	once sync.Once
}

// NewService create a service
func NewService(cc ...conf.Option) service.Service {
	opt := &conf.Options{}
	for _, c := range cc {
		c(opt)
	}
	srv := &server{
		ctx: context.Background(),
		srv: &http.Server{
			Addr: fmt.Sprintf(":%d", opt.Port),
		},
	}
	return srv
}

// SetRouter set router
func (s *server) SetRouter(h http.Handler) {
	if h == nil {
		s.srv.Handler = gin.New()
	}
	s.srv.Handler = h
}

// OnInit init service
func (s *server) OnInit() {
}

// StartUp run service
func (s *server) StartUp() error {
	s.once.Do(func() {
		if err := s.srv.ListenAndServe(); err != nil {
			return
		}
	})
	cloSign := make(chan syscall.Signal)
	<-cloSign
	return nil
}

// Stop stop service
func (s *server) Stop() {
	if s.srv.Shutdown(s.ctx) != nil {
		panic("shutdown error")
	}
}
