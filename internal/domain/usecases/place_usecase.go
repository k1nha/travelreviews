package usecase

import "github.com/k1nha/travelreviews/internal/domain/entity"

type PlaceUseCases struct {
	PlaceRepository entity.PlaceRepository
}

func NewPlaceUseCases(placeRepo entity.PlaceRepository) *PlaceUseCases {
	return &PlaceUseCases{
		PlaceRepository: placeRepo,
	}
}

func (p *PlaceUseCases) CreatePlace(name, street, city string) (*entity.Place, error) {
	place := entity.NewPlace(name, street, city)

	err := p.PlaceRepository.Save(place)
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (p *PlaceUseCases) Fetch() ([]entity.Place, error) {
	places, err := p.PlaceRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return places, nil
}
