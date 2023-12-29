package entity

import (
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
)

type Place struct {
	ID        entity.ID
	Name      string
	Street    string
	City      string
	CreatedAt time.Time
}

func NewPlace(name, street, city string) *Place {
	return &Place{
		ID:        entity.NewID(),
		Name:      name,
		Street:    street,
		City:      city,
		CreatedAt: time.Now(),
	}
}
