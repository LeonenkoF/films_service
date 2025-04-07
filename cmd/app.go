package cmd

import (
	"films_service/internal/repository"
	"fmt"
)

func Run() error {
	db, err := repository.NewPostgresDB()

	if err != nil {
		return fmt.Errorf("не удалось подключится к Бд: %v", err)
	}

	repoMovie := repository.NewMovie(db)
	repoActor := repository.NewActor(db)

	fmt.Println(repoMovie, repoActor)

	return nil
}
