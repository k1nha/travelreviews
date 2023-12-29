package placeusecase

import (
	"github.com/k1nha/travelreviews/internal/entity"
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
<<<<<<< HEAD:internal/usecases/create_place.go
	p := entity.NewPlace(
=======
	p := domain.NewPlace(
		uuid.New(),
>>>>>>> 591cd30b4615f6abf416be2709283529a6621131:internal/domain/usecases/placeusecase/create.go
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
