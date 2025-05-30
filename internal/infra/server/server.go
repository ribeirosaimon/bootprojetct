package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var _bootProjectServer *bootProjectServer

type bootProjectServer struct {
	engine *gin.Engine
	port   string
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
	_bootProjectServer = &bootProjectServer{
		port: cfg.port,
	}
	_bootProjectServer.engine = gin.Default()

	return _bootProjectServer
}

func (b *bootProjectServer) Start(f func(engine *gin.Engine)) {
	// start routes
	f(b.engine)

	if err := _bootProjectServer.engine.Run(fmt.Sprintf(":%s", b.port)); err != nil {
		panic(err)
	}
}
