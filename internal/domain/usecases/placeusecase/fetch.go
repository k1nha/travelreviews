package placeusecase

import (
	"github.com/k1nha/travelreviews/internal/domain"
)

type FetchPlacesOutput struct {
	Places []domain.Place `json:"places"`
}

type FetchPlaces struct {
	PlaceRepository domain.PlaceRepository
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
