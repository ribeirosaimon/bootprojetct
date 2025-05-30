package database

import (
	"context"
	"github.com/ribeirosaimon/bootprojetct/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunMigrations(t *testing.T) {
	ctx := context.Background()
	host, c, err := testutils.StartMysqlContainer(ctx)
	assert.Nil(t, err)
	defer c(ctx)
	connection := NewMysqlConnection(
		WithHost(host),
	)

	err = connection.RunMigrations(context.Background())
	assert.Nil(t, err)
}
