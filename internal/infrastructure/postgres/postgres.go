package postgres

import (
	"database/sql"
	"fmt"
)

func NewPostgresDB(options Options) (*sql.DB, error) {
	db, err := sql.Open("postgres", options.URL)
	if err != nil {
		return nil, fmt.Errorf("commonerror connecting to postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("commonerror pinging postgres: %w", err)
	}

	return db, nil
}

type Options struct {
	URL string
}
