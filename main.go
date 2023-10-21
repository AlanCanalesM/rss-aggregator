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

// apiConfig holds the API configuration.
type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environment variables from a .env file if present.
	loadEnvironmentVariables()

	// Set up the HTTP server.
	port := getServerPort()
	router := createRouter()
	server := createHTTPServer(port, router)

	// Create a database connection.
	db := connectToDatabase(getDBURL())
	apiCfg := apiConfig{DB: db}

	// Define and mount API routes.
	mountAPIRoutes(router, apiCfg)

	// Start the HTTP server.
	startHTTPServer(server)

	config := ScrapingConfig{
		DB:                  db,
		Concurrency:         10,
		TimeBetweenRequests: time.Minute,
	}
	// Start the scraper in the background.
	StartScraping(config)
}

// connectToDatabase creates a database connection and returns a Queries struct.
func connectToDatabase(dbURL string) *database.Queries {
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Failed connect to the db")
	}

	db := database.New(conn)
	return db
}

// loadEnvironmentVariables loads environment variables from a .env file.
func loadEnvironmentVariables() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file.")
	}
}

// getServerPort retrieves the port for the HTTP server, defaulting to 8080.
func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// createRouter creates and configures the main router for the API.
func createRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	return router
}

// createHTTPServer creates an HTTP server with the specified port and router.
func createHTTPServer(port string, router chi.Router) *http.Server {
	return &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
}

// getDBURL retrieves the database URL, exiting the program if it's not found.
func getDBURL() string {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found")
	}
	return dbURL
}

// mountAPIRoutes defines and mounts API routes on the router.
func mountAPIRoutes(router chi.Router, apiCfg apiConfig) {
	v1 := chi.NewRouter()
	// v1.Get("/signin", apiCfg.handlerSignIn)
	v1.Get("/health", handlerHealtCheck)
	v1.Get("/error", handlerError)
	v1.Get("/feeds", apiCfg.handlerGetFeeds)
	v1.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostForUser))
	v1.Get("/allPosts", apiCfg.handlerGetAllPosts)
	v1.Post("/users", apiCfg.handlerCreateUser)
	v1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1.Post("/signin", apiCfg.handlerSignin)
	v1.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))
	router.Mount("/v1", v1)
}

// startHTTPServer starts the HTTP server and logs any errors.
func startHTTPServer(server *http.Server) {
	fmt.Printf("Starting server at port: %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
	}
}
