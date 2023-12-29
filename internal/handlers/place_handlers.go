package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/usecases"
)

type PlaceHandler struct {
	PlaceDB *sql.DB
}

func NewPlaceHandler(placeDB *sql.DB) *PlaceHandler {
	return &PlaceHandler{
		PlaceDB: placeDB,
	}
}

func (h *PlaceHandler) CreatePlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var placeDTO PlaceRequest

	err := json.NewDecoder(r.Body).Decode(&placeDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = placeDTO.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := usecases.CreatePlace{
		PlaceRepository: &database.PlaceRepository{
			Db: h.PlaceDB,
		},
	}

	input := usecases.PlaceInput{
		Name:   placeDTO.Name,
		Street: placeDTO.Street,
		City:   placeDTO.City,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.PlaceDB.Close()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
