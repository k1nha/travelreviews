package database

import (
	"database/sql"
	"errors"

	"github.com/k1nha/travelreviews/internal/domain"
)

type PlaceRepository struct {
	Db *sql.DB
}

func (p *PlaceRepository) Save(place *domain.Place) error {
	_, error := p.Db.Exec("INSERT INTO place (id, name, street, city, created_at) VALUES (?, ?, ?, ?, ?)", place.ID, place.Name, place.Street, place.City, place.CreatedAt)

	if error != nil {
		return errors.New("error on create Place in Database")
	}

	return nil
}

func (p *PlaceRepository) GetById(id string) (*domain.Place, error) {
	row := p.Db.QueryRow("SELECT id, name, street, city, created_at FROM place WHERE id =?", id)

	var place domain.Place

	err := row.Scan(&place.ID, &place.Name, &place.Street, &place.City, &place.CreatedAt)

	if err != nil {
		return nil, errors.New("error on get Place from Database")
	}

	return &place, nil
}

func (p *PlaceRepository) GetAll() ([]domain.Place, error) {

	var places []domain.Place

	rows, err := p.Db.Query("SELECT * FROM places")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var place domain.Place

		if err := rows.Scan(&place.ID, &place.Name, &place.Street, &place.City, &place.CreatedAt); err != nil {
			return places, err
		}

		places = append(places, place)
	}

	if err = rows.Err(); err != nil {
		return places, err
	}

	return places, nil
}
