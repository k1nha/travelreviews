package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/k1nha/travelreviews/router"
)

func main() {
	h := router.InitializeRouter()
	godotenv.Load(".env")

	p := os.Getenv("PORT")
	if p == "" {
		log.Fatal("PORT is not found in the environment variable")
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
