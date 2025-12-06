package registry

import (
	"std-go/restserver/handler"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Handle() gin.HandlerFunc
}

func MiddlewareRegistry(names ...string) Registry {
	return func(engine *gin.Engine) {
		for _, name := range names {
			middleware := handler.GetHandler(name)
			if middleware == nil {
				continue
			}
			engine.Use(middleware.Handle)
		}
	}
}
