package router

import (
	"github.com/go-chi/chi"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/handlers"
	"github.com/k1nha/travelreviews/internal/infra/http/middlewares"
	"github.com/k1nha/travelreviews/util"
)

func initRoutes(c *chi.Mux) {
	db := database.GetDB()

	placeHandler := handlers.NewPlaceHandler(db)
	reviewHandler := handlers.NewReviewHandler(db)
	userHandler := handlers.NewUserHandler(db)

	c.HandleFunc("/healthcheck", util.HealthCheck())

	c.Route("/travel", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", reviewHandler.Create)
	})

	c.Route("/place", func(r chi.Router) {
		// r.Use(middlewares.AuthMiddleware)
		r.Post("/", placeHandler.Create)
		r.Get("/", placeHandler.Fetch)
	})

	c.Route("/user", func(r chi.Router) {
		r.Route("/tkn", func(r1 chi.Router) {
			r1.Post("/generate_token", userHandler.CreateToken)
		})
	})
}
