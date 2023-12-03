package database

import (
	"database/sql"
	"errors"

	"github.com/k1nha/travelreviews/internal/entity"
)

type PlaceRepository struct {
	Db *sql.DB
}

func (p *PlaceRepository) Save(place *entity.Place) error {
	_, error := p.Db.Exec("INSERT INTO place (place_id)")

	if error != nil {
		return errors.New("error on create Place in Database")
	}

	return nil
}
