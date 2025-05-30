package main

import (
	"context"
	"github.com/ribeirosaimon/bootprojetct/internal/infra/database"
	"github.com/ribeirosaimon/bootprojetct/internal/infra/server"
	"github.com/ribeirosaimon/bootprojetct/web"
	"os"
)

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	connection := database.NewMysqlConnection()
	connection.RunMigrations(context.Background(), "static/migrations")

	server.New(
		server.WithPort(port),
	).Start(
		web.StartHandlers,
	)
}
