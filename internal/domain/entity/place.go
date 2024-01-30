package entity

import (
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
)

type Place struct {
	ID           entity.ID `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	AveragePrice float64   `json:"average_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"update_at"`
}

func NewPlace(name, typeOf string) *Place {
	return &Place{
		ID:           entity.NewID(),
		Name:         name,
		Type:         typeOf,
		AveragePrice: 0,
		CreatedAt:    time.Now(),
		UpdateAt:     time.Now(),
	}
}

func (p *Place) Touch() {
	p.UpdateAt = time.Now()
}

type PlaceRepository interface {
	Save(place *Place) error
	GetById(id string) (*Place, error)
	GetAll() ([]Place, error)
}
