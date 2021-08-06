package main

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(rw http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
	}

	app.logger.Print("id is", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Some movie",
		Description: "Some description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(rw, http.StatusOK, movie, "movie")

	app.logger.Print("Error: ", err)
}

func (app *application) getAllMovies(rw http.ResponseWriter, r *http.Request) {

}
