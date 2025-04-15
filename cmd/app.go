package cmd

import (
	"films_service/internal/api/handlers"
	"films_service/internal/repository"
	"films_service/internal/service"
	"fmt"
	"net/http"
)

func Run() error {
	db, err := repository.NewPostgresDB()

	if err != nil {
		return fmt.Errorf("не удалось подключится к Бд: %v", err)
	}

	repoMovie := repository.NewMovie(db)
	repoActor := repository.NewActor(db)

	movieService := service.NewMoiveService(*repoActor, *repoMovie)

	h := handlers.NewMoiveServiceHandler(movieService)

	handlers.InitRoutes(h)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
