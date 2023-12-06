package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func ResponseError(w http.ResponseWriter, c int, m string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	if m != "" {
		json.NewEncoder(w).Encode(Error{
			Message: m,
		})
	}
}
