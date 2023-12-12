// Package data contains a Database interface with multiple implementations.
package data

import (
	"github.com/qdm12/go-template/internal/config/settings"
	"github.com/qdm12/go-template/internal/data/psql"
	"github.com/qdm12/log"
)

func NewPostgres(config settings.PostgresDatabase, logger log.LeveledLogger) (
	db *psql.Database, err error) {
	return psql.NewDatabase(config, logger)
}
