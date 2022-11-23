package handlers

import (
	"encoding/json"
	"fmt"
	"goMovieManagement/internals/models"
	"goMovieManagement/internals/stores"

	"github.com/gorilla/mux"

	//"goMovieManagement/models"
	"net/http"
	"strconv"
)

var Movies []models.Movie

func New(store *stores.Store) *HttpHandler {
	return &HttpHandler{store}
}

type HttpHandler struct {
	storeHandler *stores.Store
}

func (h *HttpHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to Movie management!")
	fmt.Println("endpoint Hit : homepage")
}

func (h *HttpHandler) ReturnAllMovies(w http.ResponseWriter, r *http.Request) {
	//query
	fmt.Println("Endpoint hit: get all movies")
	w.Header().Set("content-type", "application/json")
	movies := h.storeHandler.GetMoviesFromDB()
	json.NewEncoder(w).Encode(movies)
}

func (h *HttpHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("invalid id: ", err)
		return
	}

	movie := h.storeHandler.GetMovieFromDB(id)
	json.NewEncoder(w).Encode(movie)
}

//func (h *HttpHandler) deleteMovieById(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	key := vars["id"]
//	for index, movie := range Movies {
//		if movie.ID == key {
//			// delete movie here
//			Movies = append(Movies[:index], Movies[index+1:]...)
//		}
//	}
//	//json.NewEncoder(w).Encode()
//}

func (h *HttpHandler) AddNewMovie(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	//reqBody, _ := ioutil.ReadAll(r.Body)

	var movie models.Movie
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		//panic(err)
		fmt.Println("json error")
		fmt.Fprintf(w, "Incorrect JSON format")
		return
	}
	movie = h.storeHandler.AddMovieToDB(movie)
	json.NewEncoder(w).Encode(movie)
}

func (h *HttpHandler) UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie models.Movie
	err := decoder.Decode(&movie)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println("somerror")
	}
	h.storeHandler.UpdateMovieInDB(movie)
	Movies = nil
	h.storeHandler.GetMovieFromDB(movie.ID)
	json.NewEncoder(w).Encode(Movies)
	fmt.Fprintf(w, "Error updating the values")
	return
}
