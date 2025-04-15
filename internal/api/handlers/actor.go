package handlers

import (
	"encoding/json"
	"films_service/internal/models"
	"net/http"
)

func (h *MoiveServiceHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var input models.Actor
	var id int

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Gender == "" || input.Birth == "" {
		http.Error(w, "Данные не могут быть пустыми", http.StatusBadRequest)
		return
	}

	id, err := h.movieService.CreateActor(input)
	if err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)

}

func (h *MoiveServiceHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	var input models.Actor

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
		return
	}

	if input.Id == 0 {
		http.Error(w, "Данные не могут быть пустыми", http.StatusBadRequest)
		return
	}

	if err := h.movieService.DeleteActor(input.Id); err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}
}

func (h *MoiveServiceHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	var input models.Actor

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ошибка запроса", http.StatusBadRequest)
		return
	}

	if input.Id == 0 || input.Name == "" || input.Gender == "" || input.Birth == "" {
		http.Error(w, "Данные не могут быть пустыми", http.StatusBadRequest)
		return
	}

	if err := h.movieService.UpdateActor(input); err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}

}

func (h *MoiveServiceHandler) ActorsList(w http.ResponseWriter, r *http.Request) {
	movie_actors, err := h.movieService.ActorsList()
	if err != nil {
		http.Error(w, "Ошибка выполнения функции", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(movie_actors); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
