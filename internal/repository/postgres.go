package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB, error) {
	dbConnectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"postgres_films_service", "5432", "postgres", "movie", "postgres", "disable")

	db, err := sqlx.Open("postgres", dbConnectString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
