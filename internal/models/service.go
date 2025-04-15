package models

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
}

type Actor struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Birth  string `json:"birth"`
}
