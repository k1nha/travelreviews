package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/k1nha/travelreviews/internal/handlers"
	"github.com/k1nha/travelreviews/util"
)

func initRoutes(c *chi.Mux) {
	c.HandleFunc("/healthcheck", util.HealthCheck())

	c.Post("/travel", handlers.CreateReviewHandler())
	c.Post("/place", handlers.CreatePlaceHandler())
}
