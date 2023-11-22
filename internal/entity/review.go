package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID `json:"id"`
	Place       string    `json:"place"`
	Stars       int       `json:"stars"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewReview(id uuid.UUID, place string, stars int, description string) *Review {
	return &Review{
		ID:          id,
		Place:       place,
		Stars:       stars,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
