package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateCode  = errors.New("duplicate item code")
)

type Models struct {
	Item ItemModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Item: ItemModel{DB: db},
	}
}
