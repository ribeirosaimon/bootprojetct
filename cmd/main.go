package main

import (
	"github.com/ribeirosaimon/bootprojetct/internal/infra/server"
	"github.com/ribeirosaimon/bootprojetct/web"
	"os"
)

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	server.New(
		server.WithPort(port),
	).Start(
		web.StartHandlers,
	)
}
