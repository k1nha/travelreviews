package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/domain/usecases/placeusecase"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers/requests"
	"github.com/k1nha/travelreviews/internal/infra/http/response"
)

type PlaceHandler struct {
	DB *sql.DB
}

func NewPlaceHandler(db *sql.DB) *PlaceHandler {
	return &PlaceHandler{
		DB: db,
	}
}

func (p *PlaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var pla requests.CreatePlaceRequest

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

	usecase := placeusecase.NewCreatePlace(
		&database.PlaceRepository{
			Db: p.DB,
		},
		&database.CommercialInfoRepository{
			Db: p.DB,
		},
	)

	output, err := usecase.Execute(&placeusecase.Input{
		Name:    pla.Name,
		TypeOf:  pla.Typeof,
		Address: pla.Address,
		City:    pla.City,
		State:   pla.State,
		Country: pla.Country,
	})

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer p.DB.Close()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(output)
}

func (p *PlaceHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usecase := placeusecase.FetchPlaces{
		PlaceRepository: &database.PlaceRepository{
			Db: p.DB,
		},
	}

	ouput, err := usecase.Execute()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer p.DB.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ouput)
}
