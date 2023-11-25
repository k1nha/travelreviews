package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID `json:"id"`
	PlaceId     uuid.UUID `json:"place_id"`
	Stars       int       `json:"stars"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewReview(id uuid.UUID, placeId uuid.UUID, stars int, description string) *Review {
	return &Review{
		ID:          id,
		PlaceId:     placeId,
		Stars:       stars,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
