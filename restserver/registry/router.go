package registry

import "github.com/gin-gonic/gin"

type Router interface {
	Router() Registry
}

func RouterRegistry(routers ...Router) Registry {
	return func(engine *gin.Engine) {
		for _, router := range routers {
			if router == nil {
				continue
			}
			route := router.Router()
			if route == nil {
				continue
			}
			route(engine)
		}
	}
}
