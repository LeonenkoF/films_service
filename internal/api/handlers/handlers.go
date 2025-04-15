package handlers

import (
	"films_service/internal/service"
	"net/http"
)

type MoiveServiceHandler struct {
	movieService *service.MovieService
}

func NewMoiveServiceHandler(movieService *service.MovieService) *MoiveServiceHandler {
	return &MoiveServiceHandler{
		movieService: movieService,
	}
}

func InitRoutes(h *MoiveServiceHandler) {

	http.HandleFunc("/actor/create", h.CreateActor)
	http.HandleFunc("/actor/delete", h.DeleteActor)
	http.HandleFunc("/actor/update", h.UpdateActor)
	http.HandleFunc("/actor/list", h.ActorsList)

}
