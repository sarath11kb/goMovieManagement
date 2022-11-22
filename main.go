package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goMovieManagement/Models"
	"log"
	"net/http"
)

var Movies []Models.Movie

func returnAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get all movies")
	json.NewEncoder(w).Encode(Movies)

}

func returnSingleMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, movie := range Movies {
		if movie.Id == key {
			json.NewEncoder(w).Encode(movie)
		}
	}
	//json.NewEncoder(w).Encode()
}

func deleteMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for index, movie := range Movies {
		if movie.Id == key {
			// delete movie here
			Movies = append(Movies[:index], Movies[index+1:]...)
		}
	}
	//json.NewEncoder(w).Encode()
}

func updateMovieById(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie Models.Movie
	err := decoder.Decode(&movie)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println("somerror")
	}

	for index, value := range Movies {
		if value.Id == movie.Id {

			Movies[index].Name = movie.Name
			Movies[index].Released = movie.Released
			Movies[index].Plot = movie.Plot
			Movies[index].Rating = movie.Rating
			Movies[index].Released = movie.Released
			json.NewEncoder(w).Encode(Movies[index])
			return
		}
	}

	fmt.Fprintf(w, "Error updating the values")
	return
	//vars := mux.Vars(r)
	//key := vars["id"]
	//for _, movie := range Movies {
	//	if movie.Id == key {
	//		// update movie here Movies = append(Movies[:movie.Id], Movies[movie.Id+1:])
	//		//Movies[index] =
	//
	//	}
	//}
	//json.NewEncoder(w).Encode()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to Movie management!")
	fmt.Println("endpoint Hit : homepage")
}

func addNewMovie(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	//reqBody, _ := ioutil.ReadAll(r.Body)

	var movie Models.Movie
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		//panic(err)
		fmt.Println("json error")
		fmt.Fprintf(w, "Incorrect JSON format")
		return
	}
	Movies = append(Movies, movie)

	// update our global Articles array to include
	// our new Article

	json.NewEncoder(w).Encode(movie)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/movies", returnAllMovies)
	myRouter.HandleFunc("/movie", addNewMovie).Methods("POST")
	myRouter.HandleFunc("/movie/{id}", updateMovieById).Methods("PUT")
	myRouter.HandleFunc("/movie/{id}", returnSingleMovie).Methods("GET")
	myRouter.HandleFunc("/delete/movie/{id}", deleteMovieById)
	//	http.HandleFunc("/movies", returnAllMovies)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Movies = []Models.Movie{
		{
			Id:       "1",
			Name:     "Tower of god",
			Genre:    "Fantasy",
			Released: false,
		}, {
			Id:       "2",
			Name:     "Armor of go",
			Genre:    "Comedy",
			Released: true,
		},
	}
	handleRequests()
}
