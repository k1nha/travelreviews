package domain

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"nome"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type PlaceRepository interface {
	Save(place *Place) error
	GetById(id string) (*Place, error)
	GetAll() ([]Place, error)
}

func NewPlace(id uuid.UUID, name string, street string, city string) *Place {
	return &Place{
		ID:        id,
		Name:      name,
		Street:    street,
		City:      city,
		CreatedAt: time.Now(),
	}
}
