package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/domain/usecases/authusecase"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers/requests"
	"github.com/k1nha/travelreviews/internal/infra/http/response"
	"github.com/k1nha/travelreviews/util"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (h *UserHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tokenReq requests.TokenRequest

	json.NewDecoder(r.Body).Decode(&tokenReq)

	err := tokenReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	usecase := authusecase.GenerateToken{
		UserRepository: &database.UserRepository{
			Db: h.DB,
		},
		JwtAdapter: &util.JwtAdapter{},
	}

	input := authusecase.TokenInput{
		Email:    tokenReq.Email,
		Username: tokenReq.Username,
	}

	tkn, err := usecase.Execute(input)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer h.DB.Close()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(tkn)
}
