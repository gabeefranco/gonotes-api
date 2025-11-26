package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gabeefranco/gonotes-api/internal/config"
)

func NewDB(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
