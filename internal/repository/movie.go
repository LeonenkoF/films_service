package repository

import (
	"films_service/internal/entity"
	"films_service/internal/models"
	"fmt"
	"log"

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

func (a *Movie) CreateMovie(model models.Movie) (int, error) {
	var id int

	movie := &entity.Movie{
		Title:        model.Title,
		Description:  model.Description,
		Release_date: model.Release_date,
		Rating:       model.Rating,
	}

	query := "INSERT INTO movie (title, description, release_date, rating) values($1,$2,$3,$4) RETURNING id"

	rows := a.db.QueryRow(query, movie.Title, movie.Description, movie.Release_date, movie.Rating)
	if err := rows.Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (a *Movie) UpdateMovie(model models.Movie) error {

	movie := &entity.Movie{
		Id:           model.Id,
		Title:        model.Title,
		Description:  model.Description,
		Release_date: model.Release_date,
		Rating:       model.Rating,
	}

	query := "UPDATE movie SET title=$1, description=$2, release_date=$3, rating=$4 WHERE id=$5;"
	result, err := a.db.Exec(query, movie.Title, movie.Description, movie.Release_date, movie.Rating, movie.Id)

	if err != nil {
		log.Println(err, result)
		return err
	}

	return nil

}

func (a *Movie) DeleteMovie(id int) error {
	query := "DELETE FROM movie WHERE id = $1"
	result, err := a.db.Exec(query, id)

	if err != nil {
		log.Println(err, result)
		return err
	}
	return nil
}

func (a *Movie) MoviesList(sortedBy string, direct string) ([]entity.Movie, error) {
	var result []entity.Movie

	query := fmt.Sprintf("SELECT id, title, description, release_date, rating FROM movie ORDER BY %s %s", sortedBy, direct)

	fmt.Println(query)
	rows, err := a.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row entity.Movie
		if err := rows.Scan(&row.Id, &row.Title, &row.Description, &row.Release_date, &row.Rating); err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
