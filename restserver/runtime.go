package restserver

import (
	"fmt"

	"github.com/muhaobing-eng/std-go/restserver/config"
	"github.com/muhaobing-eng/std-go/restserver/lib"
	"github.com/muhaobing-eng/std-go/restserver/registry"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func Init(registries ...registry.Registry) error {
	// init dependencies
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if cfg.Database != nil {
		if err = lib.InitDatabase(cfg.Database); err != nil {
			return err
		}
	}
	if cfg.Cache != nil {
		if err = lib.InitCache(cfg.Cache); err != nil {
			return err
		}
	}

	// init rest server
	server = gin.New()
	// TODO 日志
	for _, registry := range registries {
		registry(server)
	}

	return nil
}

func Run() {
	cfg := config.Get()
	server.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
