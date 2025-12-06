package database

import (
	"github.com/muhaobing-eng/std-go/go-common/database"
	"github.com/muhaobing-eng/std-go/restserver/lib"

	"github.com/gin-gonic/gin"
)

const DatabaseHandlerKey = "database"

type DatabaseHandler struct{}

func (d *DatabaseHandler) Name() string {
	return DatabaseHandlerKey
}

func (d *DatabaseHandler) Handle(ctx *gin.Context) {
	db := lib.GetDB()
	if db != nil {
		ctx.Request = ctx.Request.WithContext(database.Context(ctx.Request.Context(), db))
	}
	ctx.Next()
}
