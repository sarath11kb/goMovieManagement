package main

import (
	"bytes"
	"encoding/json"
	"goMovieManagement/Models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddNewMovie(t *testing.T) {
	method := http.MethodPost
	endpoint := "/movie"
	testCases := []struct {
		endPoint      string
		httpMethod    string
		description   string
		expectedError error
		body          Models.Movie
	}{
		{
			endPoint:      endpoint,
			httpMethod:    method,
			expectedError: nil,
			description:   "new movie adding",
			body: Models.Movie{
				Id:       "3",
				Name:     "Police story 1",
				Genre:    "action",
				Rating:   5,
				Plot:     "",
				Released: true,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			var buff bytes.Buffer
			bodyErr := json.NewEncoder(&buff).Encode(tt.body)
			if bodyErr != nil {
				t.Errorf("error with testcase body : %s", bodyErr)
			}

			request := httptest.NewRequest(tt.httpMethod, tt.endPoint, &buff)
			w := httptest.NewRecorder()
			addNewMovie(w, request)
			res := w.Result()
			_, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Some error occured %s", err.Error())
			}
		})
	}
}
