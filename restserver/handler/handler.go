package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Name() string
	Handle(ctx *gin.Context)
}

var handlerMap = make(map[string]Handler)

func RegisterHandler(handler Handler) {
	handlerMap[handler.Name()] = handler
}

func GetHandler(name string) Handler {
	return handlerMap[name]
}
