package main

import (
	"log"
	"net/http"
	"time"

	"github.com/k1nha/travelreviews/config"
	"github.com/k1nha/travelreviews/internal/infra/database"
	"github.com/k1nha/travelreviews/internal/infra/http/router"
)

func main() {
	cfg, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	h := router.InitializeRouter()

	database.InitDatabase()

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		Handler:      h,
	}

	log.Printf("Runing on Port %v\n", cfg.Port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
