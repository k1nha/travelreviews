package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/k1nha/travelreviews/util"
)

func initRoutes(c *chi.Mux) {
	c.HandleFunc("/healthcheck", util.HealthCheck())

	c.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "its works}`))
		return
	})
}
