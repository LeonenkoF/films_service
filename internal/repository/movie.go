package repository

import "github.com/jmoiron/sqlx"

type movie struct {
	db *sqlx.DB
}

func NewMovie(db *sqlx.DB) *movie {
	return &movie{
		db: db,
	}
}

func (a *actor) CreateMovie() {

}

func (a *actor) UpdateMovie() {

}

func (a *actor) DeleteMovie() {

}
