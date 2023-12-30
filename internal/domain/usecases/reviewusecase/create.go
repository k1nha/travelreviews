package reviewusecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/k1nha/travelreviews/internal/domain/entity"
	entitypkg "github.com/k1nha/travelreviews/pkg/entity"
)

type ReviewInput struct {
	PlaceId     string
	UserId      string
	Description string
	Stars       int
}

type ReviewOutput struct {
	Review *entity.Review
}

type CreateReview struct {
	ReviewRepository entity.ReviewRepository
	PlaceRepository  entity.PlaceRepository
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

	userId, err := entitypkg.ParseID(i.UserId)

	if err != nil {
		return nil, err
	}

	review := entity.NewReview(
		placeId,
		userId,
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
