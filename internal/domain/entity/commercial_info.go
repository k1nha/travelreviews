package entity

import (
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
)

type CommercialInfo struct {
	ID        entity.ID `json:"id"`
	PlaceID   entity.ID `json:"place_id"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	Instagram string    `json:"instagram"`
	Website   string    `json:"website"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCommercialInfo(placeId entity.ID, address, city, state, country string) *CommercialInfo {
	return &CommercialInfo{
		ID:        entity.NewID(),
		PlaceID:   placeId,
		Address:   address,
		City:      city,
		State:     state,
		Country:   country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *CommercialInfo) Touch() {
	a.UpdatedAt = time.Now()
}

type CommercialInfoRepository interface {
	Save(address *CommercialInfo) error
}
