package cache

import (
	"std-go/go-common/cache"
	"std-go/restserver/lib"

	"github.com/gin-gonic/gin"
)

const CacheHandlerKey = "cache"

type CacheHandler struct{}

func (c *CacheHandler) Name() string {
	return CacheHandlerKey
}

func (c *CacheHandler) Handle(ctx *gin.Context) {
	redis := lib.GetRedis()
	if redis != nil {
		ctx.Request = ctx.Request.WithContext(cache.Context(ctx.Request.Context(), redis))
	}
	ctx.Next()
}
