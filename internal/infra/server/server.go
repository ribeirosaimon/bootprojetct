package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var _bootProjectServer *bootProjectServer

type bootProjectServer struct {
	engine *gin.Engine
}

type serverOption func(*serverConfig)

type serverConfig struct {
	port string
}

func WithPort(port string) serverOption {
	return func(cfg *serverConfig) {
		cfg.port = port
	}
}

func defaultServerConfig() *serverConfig {
	return &serverConfig{
		port: "8080",
	}
}

func New(config ...serverOption) *bootProjectServer {
	cfg := defaultServerConfig()
	for _, o := range config {
		o(cfg)
	}
	_bootProjectServer = &bootProjectServer{}
	_bootProjectServer.engine = gin.Default()
	_bootProjectServer.startHandlers()

	if err := _bootProjectServer.engine.Run(fmt.Sprintf(":%s", cfg.port)); err != nil {
		panic(err)
	}

	return _bootProjectServer
}
