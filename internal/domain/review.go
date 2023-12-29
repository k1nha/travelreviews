package domain

import (
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
)

type Review struct {
	ID          entity.ID `json:"id"`
	PlaceId     entity.ID `json:"place_id"`
	Stars       int       `json:"stars"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewReview(placeId entity.ID, stars int, description string) *Review {
	return &Review{
		ID:          entity.NewID(),
		PlaceId:     placeId,
		Stars:       stars,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
