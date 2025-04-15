package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sqlx.DB, error) {
	dbConnectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", "postgres", "movie", "qwerty", "disable")

	db, err := sqlx.Open("postgres", dbConnectString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
