package stores

import (
	"database/sql"
	"fmt"

	//"goMovieManagement/models"
	"goMovieManagement/internals/models"
	"log"
)

func New(db *sql.DB) *Store {
	return &Store{db}
}

type Store struct {
	db *sql.DB
}

func (s *Store) GetMovieFromDB(id int) models.Movie {
	//var movie models.Movie
	rows, err := s.db.Query("select id, name from movie where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		//Movies = nil
		movie := models.Movie{}
		err := rows.Scan(&movie.ID, &movie.Name)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		return movie
	}
	return models.Movie{}
}

func (s *Store) GetMoviesFromDB() []models.Movie {
	var movies []models.Movie
	rows, err := s.db.Query("select id, name from movie")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		movie := models.Movie{}
		err := rows.Scan(&movie.ID, &movie.Name)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		movies = append(movies, movie)
		fmt.Println(movies)
	}
	return movies
}

func (s *Store) UpdateMovieInDB(movie models.Movie) {
	stmt, err := s.db.Prepare("UPDATE movie set name = ?, rating = ?, plot = ? where id = ?")
	if err != nil {
		fmt.Println("error adding statement", err)
		log.Fatal(err)
	}
	res, err := stmt.Exec(movie.Name, movie.Rating, movie.Plot, movie.ID)
	if err != nil {
		fmt.Println("error executing query", err)
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error lastId")
		log.Fatal(err)

	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func (s *Store) AddMovieToDB(movie models.Movie) (models.Movie, error) {
	stmt, err := s.db.Prepare("INSERT INTO movie(name, rating, plot) VALUES(?, ?,?)")
	if err != nil {
		fmt.Println("error adding statement", err)
		log.Fatal(err)
		return models.Movie{}, err
	}

	res, err := stmt.Exec(movie.Name, movie.Rating, movie.Plot)
	if err != nil {
		fmt.Println("error executing query", err)
		log.Fatal(err)
		return models.Movie{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error lastId")
		log.Fatal(err)
		return models.Movie{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return models.Movie{}, err
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	movie.ID = int(lastId)
	return movie, nil

}
