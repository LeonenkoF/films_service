package repository

import (
	"films_service/internal/entity"
	"films_service/internal/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Actor struct {
	db *sqlx.DB
}

func NewActor(db *sqlx.DB) *Actor {
	return &Actor{
		db: db,
	}
}

func (a *Actor) CreateActor(model models.Actor) (int, error) {
	id := 0
	actor := &entity.Actor{
		Name:   model.Name,
		Gender: model.Gender,
		Birth:  model.Birth,
	}

	query := "INSERT INTO actors (name,gender,birth) values($1,$2,$3) RETURNING id"
	rows := a.db.QueryRow(query, actor.Name, actor.Gender, actor.Birth)
	if err := rows.Scan(&id); err != nil {
		log.Println(err)
		return 0, fmt.Errorf("ошибка обращения к db:", err)
	}
	return id, nil
}

func (a *Actor) UpdateActor(model models.Actor) error {

	actor := &entity.Actor{
		Id:     model.Id,
		Name:   model.Name,
		Gender: model.Gender,
		Birth:  model.Birth,
	}

	query := "UPDATE actors SET name=$1, gender=$2, birth=$3 WHERE id=$4;"
	result, err := a.db.Exec(query, actor.Name, actor.Gender, actor.Birth, actor.Id)

	if err != nil {
		log.Println(err, result)
		return err
	}

	return nil
}

func (a *Actor) DeleteActor(id int) error {
	query := "DELETE FROM actors WHERE id = $1"
	result, err := a.db.Exec(query, id)

	if err != nil {
		log.Println(err, result)
		return err
	}
	return nil
}

func (a *Actor) ActorsList() ([]entity.Movie_actor, error) {
	var result []entity.Movie_actor
	query := `SELECT
    a.id AS actor_id,
    a.name AS actor_name,
    a.gender,
    a.birth,
    m.title AS movie_title
FROM
    actors a
JOIN
    movie_actors ma ON a.id = ma.actor_id
JOIN
    movie m ON m.id = ma.movie_id
ORDER BY
    a.name, m.release_date;`

	rows, err := a.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row entity.Movie_actor
		if err := rows.Scan(&row.Id, &row.Name, &row.Gender, &row.Birth, &row.Movie_title); err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
