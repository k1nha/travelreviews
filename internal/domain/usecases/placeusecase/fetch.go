package placeusecase

import "github.com/k1nha/travelreviews/internal/domain/entity"

type FetchPlacesOutput struct {
	Places []entity.Place `json:"places"`
}

type FetchPlaces struct {
	PlaceRepository entity.PlaceRepository
}

func (f *FetchPlaces) Execute() (*FetchPlacesOutput, error) {
	places, err := f.PlaceRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return &FetchPlacesOutput{
		Places: places,
	}, nil
}
