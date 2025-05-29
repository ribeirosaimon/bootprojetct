package main

import (
	"github.com/ribeirosaimon/bootprojetct/internal/infra/server"
	"os"
)

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	server.New(server.WithPort(port))
}
