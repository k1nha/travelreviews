package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/http/response"
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
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = rev.Validate()

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		// TODO: Find if place exists
		usecase := usecases.CreateReview{
			ReviewRepository: &database.ReviewRepository{
				Db: db,
			},
			PlaceRepository: &database.PlaceRepository{
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
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		db.Close()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(output)
	}
}
