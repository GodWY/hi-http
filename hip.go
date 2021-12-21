package hip

import (
	"github.com/GodWY/hip/app"
	"github.com/GodWY/hip/service"
)

func NewService(cc *app.Options) service.Service {
	return app.NewHttp(cc)
}
