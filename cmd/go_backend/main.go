package main

import (
	"flc_backend/internal/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port string is not provided.\n")
	}

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: routes.NewRouter(),
	}

	log.Printf("listening at port: %s", portString)

	srv.ListenAndServe()

}
