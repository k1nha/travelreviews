package placeusecase

import "github.com/k1nha/travelreviews/internal/domain/entity"

type output struct {
	Place *entity.Place
}

type CreatePlace struct {
	placeRepository       entity.PlaceRepository
	commercialInfoRepository entity.CommercialInfoRepository
}

func NewCreatePlace(placeRepo entity.PlaceRepository, commercialInfoRepo entity.CommercialInfoRepository) *CreatePlace {
	return &CreatePlace{
		placeRepository:       placeRepo,
		commercialInfoRepository: commercialInfoRepo,
	}
}

func (c *CreatePlace) Execute(name, typeof string) (*output, error) {
	p := entity.NewPlace(name, typeof)

	err := c.placeRepository.Save(p)

	if err != nil {
		return nil, err
	}

	return &output{
		Place: p,
	}, nil
}
