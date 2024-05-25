package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(dataSourceName string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{db: db}, nil
}

// Implement methods for PostgresDB
