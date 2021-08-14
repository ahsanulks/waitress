package config_test

import (
	"os"
	"testing"

	"github.com/ahsanulks/waitress/config"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
)

func TestNewPostgresConf(t *testing.T) {
	gotenv.Load("../.env")
	conf := config.NewPostgresConf()

	assert.NotEmpty(t, conf.Database)
	assert.NotEmpty(t, conf.Host)
	assert.NotEmpty(t, conf.Port)
	assert.NotEmpty(t, conf.Username)
	assert.NotEmpty(t, conf.Password)

	assert.Equal(t, os.Getenv("POSTGRESQL_USERNAME"), conf.Username)
	assert.Equal(t, os.Getenv("POSTGRESQL_PASSWORD"), conf.Password)
	assert.Equal(t, os.Getenv("POSTGRESQL_HOST"), conf.Host)
	assert.Equal(t, os.Getenv("POSTGRESQL_DATABASE"), conf.Database)
	assert.Equal(t, os.Getenv("POSTGRESQL_PORT"), conf.Port)

	os.Setenv("POSTGRES_TIMEOUT", "20")

	conf = config.NewPostgresConf()
	assert.Equal(t, "20", conf.Timeout)
}

type fakeLog struct{}

func (fl fakeLog) Store(err error, message string, options map[string]interface{}) {}

func TestNewPostgresConn(t *testing.T) {
	gotenv.Load("../.env")
	conf := config.NewPostgresConf()
	fl := new(fakeLog)

	repo, adapter := config.NewPostgresConn(conf, fl)
	assert.Implements(t, (*rel.Repository)(nil), repo)
	assert.IsType(t, &postgres.Adapter{}, adapter)
}
