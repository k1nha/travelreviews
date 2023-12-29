package placehandler

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/domain/usecases/placeusecase"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/response"
)

func CreatePlaceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var pla CreatePlaceRequest

		db := database.GetDB()

		err := json.NewDecoder(r.Body).Decode(&pla)

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = pla.Validate()

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		usecase := placeusecase.CreatePlace{
			PlaceRepository: &database.PlaceRepository{
				Db: db,
			},
		}

		input := placeusecase.PlaceInput{
			Name:   pla.Name,
			Street: pla.Street,
			City:   pla.City,
		}

		output, err := usecase.Execute(input)

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		db.Close()

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(output)
	}
}