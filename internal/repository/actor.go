package repository

import "github.com/jmoiron/sqlx"

type actor struct {
	db *sqlx.DB
}

func NewActor(db *sqlx.DB) *actor {
	return &actor{
		db: db,
	}
}

func (a *actor) CreateActor() {

}

func (a *actor) UpdateActor() {

}

func (a *actor) DeleteActor() {

}
