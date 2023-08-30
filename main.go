package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	v1 := chi.NewRouter()
	v1.Get("/health", handlerHealtCheck)

	router.Mount("/v1", v1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Starting server at port: %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Printf("Error starting server %v", err)
	}

}