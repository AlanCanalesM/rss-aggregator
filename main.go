package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// This is the main function of the application
	// It will start the server and the scraper
	// It will also load the environment variables
	// and create the database connection

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Failed connect to the db")
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	startScraping(db, 10, time.Minute)

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
	v1.Get("/error", handlerError)
	v1.Get("/feeds", apiCfg.handlerGetFeeds)
	v1.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostForUser))

	v1.Post("/users", apiCfg.handlerCreateUser)
	v1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))

	v1.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))

	router.Mount("/v1", v1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Starting server at port: %v", portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Printf("Error starting server %v", err)
	}

}
