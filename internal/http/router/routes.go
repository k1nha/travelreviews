package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/k1nha/travelreviews/internal/handlers"
	"github.com/k1nha/travelreviews/internal/http/middlewares"
	"github.com/k1nha/travelreviews/util"
)

func initRoutes(c *chi.Mux) {
	c.HandleFunc("/healthcheck", util.HealthCheck())

	c.Route("/travel", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", handlers.CreateReviewHandler())
	})

	c.Route("/place", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", handlers.CreatePlaceHandler())
	})

	c.Route("/auth", func(r chi.Router) {
		r.Post("/token", handlers.CreateTokenHandler())
	})
}
