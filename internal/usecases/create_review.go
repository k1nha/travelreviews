package usecases

import (
	"github.com/google/uuid"
	"github.com/k1nha/travelreviews/internal/entity"
)

type ReviewInput struct {
	PlaceId string
}

type ReviewOutput struct {
}

type CreateReview struct {
	ReviewRepository entity.ReviewRepository
}

func (c *CreateReview) Execute(i ReviewInput) (*ReviewOutput, error) {
	
	review := entity.NewReview(
		ID: uuid.New(),
		PlaceId: uuid.Parse(i.PlaceId),
		
	)

	err := c.ReviewRepository.Save(review)

	if err != nil {
		return nil, err

	}

	return &ReviewOutput{}, err
}
