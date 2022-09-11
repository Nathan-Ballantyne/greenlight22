package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Nathan-Ballantyne/greenlight22/internal/data"
	"github.com/Nathan-Ballantyne/greenlight22/internal/validator"
)

const (
	errMustBeProvided = "must be provided"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(strings.TrimSpace(input.Title) != "", "title", errMustBeProvided)
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(input.Year != 0, "year", errMustBeProvided)
	v.Check(input.Year >= 1888, "year", "must be greater then 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(input.Runtime != 0, "runtime", errMustBeProvided)
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(input.Genres != nil, "genres", errMustBeProvided)
	v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
