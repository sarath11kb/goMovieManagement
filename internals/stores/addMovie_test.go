package stores

import (
	"fmt"
	"goMovieManagement/internals/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

//func TestGetMovieById(t *testing.T) {
//	method := http.MethodGet
//	endPoint := "/movie"
//	movieNotFoundError := "Movie not found"
//	testCases := []struct {
//		id       string `json:"id"`
//		testName string
//		err      error
//	}{
//		{
//			id:       "1",
//			testName: "valid id",
//			err:      nil,
//		}, {
//			id:       "2",
//			testName: "valid id",
//			err:      nil,
//		}, {
//			id:       "3",
//			testName: "invalid id",
//			err:      errors.New(movieNotFoundError),
//		},
//	}
//	for _, tt := range testCases {
//		target := endPoint + "/" + tt.id
//		movieRequest := httptest.NewRequest(method, target, nil)
//		writer := httptest.NewRecorder()
//		t.Run(tt.testName, func(t *testing.T) {
//			getMovieById(writer, movieRequest)
//			response := writer.Result()
//			_, respError := io.ReadAll(response.Body)
//			if respError != nil {
//				t.Errorf("some error occured : %s", respError)
//			}
//			//else {
//			//	response.Body.
//			//}
//			//} else {
//			//	fmt.Printf("body data is : %s", )
//			//}
//		})
//	}
//}
//
//func TestAddNewMovie(t *testing.T) {
//	method := http.MethodPost
//	endpoint := "/movie"
//	testCases := []struct {
//		endPoint      string
//		httpMethod    string
//		description   string
//		expectedError error
//		body          models.Movie
//	}{
//		{
//			endPoint:      endpoint,
//			httpMethod:    method,
//			expectedError: nil,
//			description:   "new movie adding",
//			body: models.Movie{
//				ID:       "3",
//				Name:     "Police story 1",
//				Genre:    "action",
//				Rating:   5,
//				Plot:     "",
//				Released: true,
//			},
//		},
//	}
//
//	for _, tt := range testCases {
//		t.Run(tt.description, func(t *testing.T) {
//			var buff bytes.Buffer
//			bodyErr := json.NewEncoder(&buff).Encode(tt.body)
//			if bodyErr != nil {
//				t.Errorf("error with testcase body : %s", bodyErr)
//			}
//
//			request := httptest.NewRequest(tt.httpMethod, tt.endPoint, &buff)
//			w := httptest.NewRecorder()
//			addNewMovie(w, request)
//			res := w.Result()
//			_, err := io.ReadAll(res.Body)
//			if err != nil {
//				t.Errorf("Some error occured %s", err.Error())
//			}
//		})
//	}
//}

func TestAddMovie(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	mockStoreHandler := New(db)
	movieVars := struct {
		validName   string
		validRating int
		validPlot   string
	}{
		validName:   "DragonBall Z",
		validRating: 4,
		validPlot:   "The adventures of Son goku ",
	}
	var (
		testCases = []struct {
			testName     string
			movie        models.Movie
			expectedExec *sqlmock.ExpectedExec
		}{
			{
				testName: "valid movie",
				movie: models.Movie{
					Name:     movieVars.validName,
					Genre:    "",
					Rating:   movieVars.validRating,
					Plot:     movieVars.validPlot,
					Released: false,
				},

				expectedExec: mock.ExpectPrepare("INSERT INTO movie(name, rating, plot) VALUES(?, ?,?)").ExpectExec().WithArgs(movieVars.validName, movieVars.validRating, movieVars.validPlot).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		}
	)
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := mockStoreHandler.AddMovieToDB(tt.movie)
			if err != nil {
				t.Errorf(err.Error())
				fmt.Println(err)

			}
		})
	}
}
