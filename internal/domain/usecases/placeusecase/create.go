package placeusecase

import "github.com/k1nha/travelreviews/internal/domain/entity"

type PlaceInput struct {
	Name   string
	Street string
	City   string
}

type PlaceOutput struct {
	Place *entity.Place
}

type CreatePlace struct {
	PlaceRepository entity.PlaceRepository
}

func (c *CreatePlace) Execute(i PlaceInput) (*PlaceOutput, error) {
	p := entity.NewPlace(
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
