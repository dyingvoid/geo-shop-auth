package postgres

import (
	"database/sql"
	"fmt"
)

func NewPostgresDB(options PostgresOptions) (*sql.DB, error) {
	db, err := sql.Open("postgres", options.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging postgres: %w", err)
	}

	return db, nil
}

type PostgresOptions struct {
	ConnectionString string
}
