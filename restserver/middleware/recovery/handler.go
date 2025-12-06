package recovery

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

const RecoveryHandlerKey = "recovery"

type Recovery struct{}

func (r *Recovery) Name() string {
	return RecoveryHandlerKey
}

func (r *Recovery) Handle(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			stack := debug.Stack()
			log.Printf("[Recovery] panic recover: %+v\n%s", err, string(stack))

			ctx.JSON(http.StatusInternalServerError, gin.H{"server panic": err})
		}
	}()
	ctx.Next()
}
