package tokenhandler

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/domain/usecases/authusecase"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/response"
	"github.com/k1nha/travelreviews/util"
)

func CreateTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var tokenReq TokenRequest

		json.NewDecoder(r.Body).Decode(&tokenReq)

		err := tokenReq.Validate()

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		db := database.GetDB()

		usecase := authusecase.GenerateToken{
			UserRepository: &database.UserRepository{
				Db: db,
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

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(tkn)
	}
}
