package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Credentials holds the user login credentials.
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims defines the JWT claims structure.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// handlerSignin handles user authentication and token creation.
func (apiCfg *apiConfig) handlerSignin(w http.ResponseWriter, r *http.Request) {
	jwtKey := []byte("my_secret_key")
	creds := Credentials{}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user := apiCfg.GetUser(w, r)

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error signing JWT: %v", err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// handlerCreateUser handles the creation of a new user.
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Username:  params.Username,
		Password:  params.Password,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

// handlerGetUser handles the retrieval of a user.
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userID")
	userIdUUID, err := uuid.Parse(userIdStr)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing UUID: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), userIdUUID)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

// handlerGetPostForUser retrieves posts for a specific user.
func (apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userID")
	userIdUUID, err := uuid.Parse(userIdStr)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing UUID: %v", err))
		return
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: userIdUUID,
		Limit:  10,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Error fetching user posts")
		return
	}

	responseWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}

// handlerGetAllPosts retrieves all posts.
func (apiCfg *apiConfig) handlerGetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := apiCfg.DB.GetAllPosts(r.Context())

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Error fetching all posts")
		return
	}

	responseWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
