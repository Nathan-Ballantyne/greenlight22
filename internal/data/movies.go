package data

import (
	"database/sql"
	"strings"
	"time"

	"github.com/Nathan-Ballantyne/greenlight22/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // release year
	Runtime   Runtime   `json:"runtime,omitempty"` // Movie runtime in minutes
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // The version number starts at 1 and will be incremented each time the movie information is updated
}

// struct that wraps an sql.DB connection pool
type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(strings.TrimSpace(movie.Title) != "", "title", validator.ErrMustBeProvided)
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", validator.ErrMustBeProvided)
	v.Check(movie.Year >= 1888, "year", "must be greater then 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", validator.ErrMustBeProvided)
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", validator.ErrMustBeProvided)
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
