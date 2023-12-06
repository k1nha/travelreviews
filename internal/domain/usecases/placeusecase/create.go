package placeusecase

import (
	"github.com/google/uuid"
	"github.com/k1nha/travelreviews/internal/domain"
)

type PlaceInput struct {
	Name   string
	Street string
	City   string
}

type PlaceOutput struct {
	Place *domain.Place
}

type CreatePlace struct {
	PlaceRepository domain.PlaceRepository
}

func (c *CreatePlace) Execute(i PlaceInput) (*PlaceOutput, error) {
	p := domain.NewPlace(
		uuid.New(),
		i.Name,
		i.Street,
		i.City,
	)

	err := c.PlaceRepository.Save(p)

	if err != nil {
		return nil, err
	}

	return &PlaceOutput{
		Place: p,
	}, nil
}
