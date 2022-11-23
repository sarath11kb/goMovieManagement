package main

import (
	"database/sql"
	"fmt"
	"goMovieManagement/internals/handlers"
	"goMovieManagement/internals/stores"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var err error
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/movie")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("database connection failed;", err)
	} else {
		fmt.Println("database connected . Ping successful")
	}
	storeHandler := stores.New(db)
	h := handlers.New(storeHandler)

	//Movies = []models.Movie{
	//	{
	//		ID:       "1",
	//		Name:     "Tower of god",
	//		Genre:    "Fantasy",
	//		Released: false,
	//	}, {
	//		ID:       "2",
	//		Name:     "Armor of go",
	//		Genre:    "Comedy",
	//		Released: true,
	//	},
	//}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", h.HomePage)
	myRouter.HandleFunc("/movies", h.ReturnAllMovies)
	myRouter.HandleFunc("/movie", h.AddNewMovie).Methods("POST")
	myRouter.HandleFunc("/movie/{id}", h.UpdateMovieById).Methods("PUT")
	myRouter.HandleFunc("/movie/{id}", h.GetMovieById).Methods("GET")
	//myRouter.HandleFunc("/delete/movie/{id}", h.deleteMovieById)
	//	http.HandleFunc("/movies", returnAllMovies)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
