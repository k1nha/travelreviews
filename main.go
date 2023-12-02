package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/k1nha/travelreviews/internal/http/router"
	"github.com/k1nha/travelreviews/internal/infra/database"
)

func main() {
	godotenv.Load(".env")

	h := router.InitializeRouter()

	database.InitDatabase()

	p := os.Getenv("PORT")
	if p == "" {
		log.Fatal("PORT not found in the environment variable")
	}

	srv := http.Server{
		Addr:         ":" + p,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		Handler:      h,
	}

	log.Printf("Runing on Port %v\n", p)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
