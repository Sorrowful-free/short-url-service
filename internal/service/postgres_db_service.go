package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDBService struct {
	db *sql.DB
}

func NewPostgresDBService(databaseDSN string) (DBService, error) {
	db, err := sql.Open("postgres", databaseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &PostgresDBService{
		db: db,
	}, nil
}

func (s *PostgresDBService) Ping() error {
	return s.db.Ping()
}
