package entity

import (
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
)

type ContactInfo struct {
	ID        entity.ID `json:"id"`
	UserID    entity.ID `json:"user_id"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContactInfo(city, state, country string, userId entity.ID) (*ContactInfo, error) {
	return &ContactInfo{
		ID:        entity.NewID(),
		UserID:    userId,
		City:      city,
		State:     state,
		Country:   country,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}, nil
}

type ContactInfoRepository interface {
	Create(contactInfo *ContactInfo) error
	// Update() error
}
