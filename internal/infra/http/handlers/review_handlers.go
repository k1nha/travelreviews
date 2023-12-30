package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/domain/usecases/reviewusecase"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers/requests"
	"github.com/k1nha/travelreviews/internal/infra/http/response"
)

type ReviewHandler struct {
	DB *sql.DB
}

func NewReviewHandler(db *sql.DB) *ReviewHandler {
	return &ReviewHandler{
		DB: db,
	}
}

func (h *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var rev requests.ReviewRequest

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

	usecase := reviewusecase.CreateReview{
		ReviewRepository: &database.ReviewRepository{
			Db: h.DB,
		},
		PlaceRepository: &database.PlaceRepository{
			Db: h.DB,
		},
	}

	input := reviewusecase.ReviewInput{
		PlaceId:     rev.PlaceId,
		Description: rev.Description,
		Stars:       rev.Stars,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer h.DB.Close()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
