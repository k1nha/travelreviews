package entity

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID        uuid.UUID
	Name      string
	Street    string
	City      string
	CreatedAt time.Time
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
