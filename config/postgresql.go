package config

import (
	"fmt"
	"os"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Database string
	Port     string
	Timeout  string
}

// NewPostgresConf initiate config for connection postgresql.
func NewPostgresConf() *PostgresConfig {
	var timeout string
	if timeout = os.Getenv("POSTGRES_TIMEOUT"); timeout == "" {
		// set default timeout on 30s
		timeout = "30"
	}
	return &PostgresConfig{
		Username: os.Getenv("POSTGRESQL_USERNAME"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Host:     os.Getenv("POSTGRESQL_HOST"),
		Database: os.Getenv("POSTGRESQL_DATABASE"),
		Port:     os.Getenv("POSTGRESQL_PORT"),
		Timeout:  timeout,
	}
}

// NewPostgresConn initiate postgres connection.
func NewPostgresConn(conf *PostgresConfig) (rel.Repository, *postgres.Adapter) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Timeout)
	adapter, err := postgres.Open(dsn)
	if err != nil {
		panic(err)
	}

	// initialize REL's repo.
	repo := rel.New(adapter)
	return repo, adapter
}
