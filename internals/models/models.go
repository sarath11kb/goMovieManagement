package models

import (
	"database/sql"
	"net/http"
)

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Rating   int    `json:"rating"`
	Plot     string `json:"plot"`
	Released bool   `json:"released"`
}

type Store struct {
	db *sql.DB
}

type HttpHandlers interface {
	updateMovieById(w http.ResponseWriter, r *http.Request)
	returnAllMovies(w http.ResponseWriter, r *http.Request)
	deleteMovieById(w http.ResponseWriter, r *http.Request)
	addNewMovie(w http.ResponseWriter, r *http.Request)
}
