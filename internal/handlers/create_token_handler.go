package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/k1nha/travelreviews/internal/http/response"
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

		// TODO: Check if Email/Username exists
		tkn, err := util.CreateToken(util.InputToken{
			Email:    tokenReq.Email,
			Username: tokenReq.Username,
		})

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)

		res := struct {
			Token string `json:"token"`
		}{
			Token: tkn,
		}

		json.NewEncoder(w).Encode(res)
	}
}
