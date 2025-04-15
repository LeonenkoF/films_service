package entity

type Movie struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Release_date string `json"release_date"`
	Rating       int    `json:"rating"`
}

type Actor struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Birth  string `json:"birth"`
}

type Movie_actor struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Birth       string `json:"birth"`
	Movie_title string `json:"movie_title"`
}
