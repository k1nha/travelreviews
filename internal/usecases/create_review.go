package usecases

import (
	"github.com/google/uuid"
	"github.com/k1nha/travelreviews/internal/entity"
)

type ReviewInput struct {
	PlaceId     string
	Description string
	Stars       int
}

type ReviewOutput struct {
	Review *entity.Review
}

type CreateReview struct {
	ReviewRepository entity.ReviewRepository
}

func (c *CreateReview) Execute(i ReviewInput) (*ReviewOutput, error) {
	placeId, err := uuid.Parse(i.PlaceId)
	if err != nil {
		return nil, err
	}

	review := entity.NewReview(
		uuid.New(),
		placeId,
		i.Stars, i.Description,
	)

	err = c.ReviewRepository.Save(review)

	if err != nil {
		return nil, err
	}

	return &ReviewOutput{
		Review: review,
	}, err
}
