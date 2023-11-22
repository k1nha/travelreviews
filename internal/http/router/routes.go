package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/k1nha/travelreviews/util"
)

func initRoutes(c *chi.Mux) {
	c.HandleFunc("/healthcheck", util.HealthCheck())
}
