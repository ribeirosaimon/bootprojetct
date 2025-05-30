package testutils

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

func StartMysqlContainer(ctx context.Context) (endpoint string, close func(context.Context), err error) {
	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "rootpassword",
			"MYSQL_DATABASE":      "shortener",
			"MYSQL_USER":          "myuser",
			"MYSQL_PASSWORD":      "mypass",
		},
		WaitingFor: wait.ForListeningPort("3306/tcp"),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	close = func(c context.Context) {
		mysqlC.Terminate(c)
	}

	endpoint, err = mysqlC.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	return
}
