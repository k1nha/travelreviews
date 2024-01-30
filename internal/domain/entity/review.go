package entity

import (
	"errors"
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
)

type Review struct {
	ID             entity.ID `json:"id"`
	PlaceId        entity.ID `json:"place_id"`
	UserId         entity.ID `json:"user_id"`
	Stars          float64   `json:"stars"`
	Description    string    `json:"description"`
	Recommendation float64   `json:"recommation"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewReview(placeId, userId entity.ID, stars, recommation float64, description string) *Review {
	return &Review{
		ID:             entity.NewID(),
		PlaceId:        placeId,
		Stars:          stars,
		UserId:         userId,
		Recommendation: recommation,
		Description:    description,
		CreatedAt:      time.Now(),
	}
}

func (r *Review) Validate() error {
	if r.Stars < 0 || r.Stars > 5 {
		return errors.New("avaliação precisa ser entre 0 a 5 estrelas")
	}
	return nil
}

type ReviewRepository interface {
	Save(review *Review) error
	GetAllByUser(userId string) ([]Review, error)
	GetAllByPlaceId(placeId string) ([]Review, error)
}
