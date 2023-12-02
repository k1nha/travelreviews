package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/usecases"
)

func CreateReviewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var rev ReviewRequest
		db := database.GetDB()

		err := json.NewDecoder(r.Body).Decode(&rev)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		usecase := usecases.CreateReview{
			ReviewRepository: &database.ReviewRepository{
				Db: db,
			},
		}

		input := usecases.ReviewInput{
			PlaceId:     rev.PlaceId,
			Description: rev.Description,
			Stars:       rev.Stars,
		}

		output, err := usecase.Execute(input)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(output)
	}
}
