package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/usecases"
)

func CreatePlaceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var pla PlaceRequest

		db := database.GetDB()

		err := json.NewDecoder(r.Body).Decode(&pla)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = pla.Validate()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		usecase := usecases.CreatePlace{
			PlaceRepository: &database.PlaceRepository{
				Db: db,
			},
		}

		input := usecases.PlaceInput{
			Name:   pla.Name,
			Street: pla.Street,
			City:   pla.City,
		}

		output, err := usecase.Execute(input)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db.Close()

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(output)
	}
}
