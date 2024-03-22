package placeusecase

import "github.com/k1nha/travelreviews/internal/domain/entity"

type Input struct {
	Name    string
	TypeOf  string
	Address string
	City    string
	State   string
	Country string
}

type output struct {
	Place *entity.Place
}

type CreatePlace struct {
	placeRepository          entity.PlaceRepository
	commercialInfoRepository entity.CommercialInfoRepository
}

func NewCreatePlace(placeRepo entity.PlaceRepository, commercialInfoRepo entity.CommercialInfoRepository) *CreatePlace {
	return &CreatePlace{
		placeRepository:          placeRepo,
		commercialInfoRepository: commercialInfoRepo,
	}
}

func (cp *CreatePlace) Execute(input *Input) (*output, error) {
	p := entity.NewPlace(input.Name, input.TypeOf)
	c := entity.NewCommercialInfo(p.ID, input.Address, input.City, input.State, input.Country)

	err := cp.placeRepository.Save(p)
	if err != nil {
		return nil, err
	}

	err = cp.commercialInfoRepository.Save(c)
	if err != nil {
		return nil, err
	}

	return &output{
		Place: p,
	}, nil
}
