package data

import (
	"database/sql"
	"errors"
)

// Custom ErrRecordNotFound error. This will be returned from the Get() method when
// looking up a movie that doesn't exist in the database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
