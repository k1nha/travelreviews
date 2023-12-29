package reviewusecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/k1nha/travelreviews/internal/domain"
)

type ReviewInput struct {
	PlaceId     string
	Description string
	Stars       int
}

type ReviewOutput struct {
	Review *domain.Review
}

type CreateReview struct {
	ReviewRepository domain.ReviewRepository
	PlaceRepository  domain.PlaceRepository
}

func (c *CreateReview) Execute(i ReviewInput) (*ReviewOutput, error) {
	placeId, err := uuid.Parse(i.PlaceId)
	if err != nil {
		return nil, err
	}

	p, err := c.PlaceRepository.GetById(i.PlaceId)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, errors.New("place not found")
	}

	review := entity.NewReview(
		placeId,
		i.Stars,
		i.Description,
	)

	err = c.ReviewRepository.Save(review)

	if err != nil {
		return nil, err
	}

	return &ReviewOutput{
		Review: review,
	}, err
}
