package handlers

import (
	"encoding/json"
	"films_service/internal/models"
	"fmt"
	"log"
	"net/http"
)

func (h *MoiveServiceHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var input models.Movie
	var id int
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err)
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
		return
	}

	if input.Title == "" || input.Rating > 10 || input.Rating < 0 || input.Release_date == "" {
		http.Error(w, "Данные не могут быть пустыми", http.StatusBadRequest)
		return
	}

	id, err := h.movieService.CreateMovie(input)
	if err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)

}

func (h *MoiveServiceHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var input models.Movie
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err)
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
		return
	}

	if input.Title == "" || input.Rating > 10 || input.Rating < 0 || input.Release_date == "" || input.Id == 0 {
		http.Error(w, "Данные не могут быть пустыми", http.StatusBadRequest)
		return
	}

	err := h.movieService.UpdateMovie(input)
	if err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}
}

func (h *MoiveServiceHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	var input models.Movie

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
		return
	}

	if input.Id == 0 {
		http.Error(w, "Данные не могут быть пустыми", http.StatusBadRequest)
		return
	}

	if err := h.movieService.DeleteMovie(input.Id); err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}
}

func (h *MoiveServiceHandler) MoviesList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	sortedBy := query.Get("sortedBy")
	direct := query.Get("direct")

	fmt.Println(direct)

	if sortedBy == "" || direct == "" || direct != "DESC" || direct != "ASC" {
		sortedBy = "rating"
		direct = "DESC"
	}

	movies, err := h.movieService.MoviesList(sortedBy, direct)
	if err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
