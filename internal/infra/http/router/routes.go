package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers/placehandler"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers/reviewhandler"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers/tokenhandler"
	"github.com/k1nha/travelreviews/internal/infra/http/middlewares"
	"github.com/k1nha/travelreviews/util"
)

func initRoutes(c *chi.Mux) {
	c.HandleFunc("/healthcheck", util.HealthCheck())

	c.Route("/travel", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", reviewhandler.CreateReviewHandler())
	})

	c.Route("/place", func(r chi.Router) {
		// r.Use(middlewares.AuthMiddleware)
		r.Post("/", placehandler.CreatePlaceHandler())
		r.Get("/", placehandler.FetchPlaceHandler())
	})

	c.Route("/auth", func(r chi.Router) {
		r.Post("/token", tokenhandler.CreateTokenHandler())
	})
}
