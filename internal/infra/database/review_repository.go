package database

import (
	"database/sql"
	"errors"

	"github.com/k1nha/travelreviews/internal/entity"
)

type ReviewRepository struct {
	Db *sql.DB
}

func (r *ReviewRepository) Save(review entity.Review) error {
	_, err := r.Db.Exec("INSERT INTO ")

	if err != nil {
		return errors.New("errors")
	}

	return nil
}
