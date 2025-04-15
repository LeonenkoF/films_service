package repository

import (
	"github.com/jmoiron/sqlx"
)

type Movie struct {
	db *sqlx.DB
}

func NewMovie(db *sqlx.DB) *Movie {
	return &Movie{
		db: db,
	}
}

func (a *Movie) CreateMovie() {
}

func (a *Movie) UpdateMovie() {

}

func (a *Movie) DeleteMovie() {

}
