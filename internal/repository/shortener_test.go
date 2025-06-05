package repository

import (
	"context"
	"github.com/ribeirosaimon/bootprojetct/internal/infra/database"
	"github.com/ribeirosaimon/bootprojetct/internal/model/entity"
	"github.com/ribeirosaimon/bootprojetct/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortener(t *testing.T) {
	var (
		localDatabase = true
		host          = "localhost:3306"
		err           error
		closeFunc     func(context.Context)
		ctx           = context.Background()
	)

	if localDatabase == false {
		host, closeFunc, err = testutils.StartMysqlContainer(ctx)
		defer closeFunc(ctx)
		assert.Nil(t, err)
	}

	connection := database.NewMysqlConnection(
		database.WithHost(host),
	)

	err = connection.RunMigrations(context.Background())

	repository := NewShortener()

	allUrls, err := repository.GetAllShortenerUrl(ctx)
	assert.Nil(t, err)
	assert.NotEmpty(t, allUrls)

	for _, v := range []struct {
		name    string
		auxFunc func()
	}{
		{
			name: "Need to return same url",
			auxFunc: func() {
				var result *entity.ShortenerUrl
				result, err = repository.GetShortenerUrlById(ctx, allUrls[0].ID)
				assert.Nil(t, err)
				assert.Equal(t, result.ID, allUrls[0].ID)
				assert.Equal(t, result.ShortCode, allUrls[0].ShortCode)
				assert.Equal(t, result.CreatedBy, allUrls[0].CreatedBy)
			},
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			v.auxFunc()
		})
	}
}
