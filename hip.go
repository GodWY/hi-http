package hip

import (
	"github.com/GodWY/hip/app"
	"github.com/GodWY/hip/service"
)

// NewService 创建一个gin服务
func NewService(cc *app.Options) service.Service {
	return app.NewHTTP(cc)
}
