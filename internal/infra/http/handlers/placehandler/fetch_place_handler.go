package placehandler

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/domain/usecases/placeusecase"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/response"
)

func FetchPlaceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		db := database.GetDB()

		usecase := placeusecase.FetchPlaces{
			PlaceRepository: &database.PlaceRepository{
				Db: db,
			},
		}

		ouput, err := usecase.Execute()

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ouput)
	}
}
