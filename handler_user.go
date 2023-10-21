package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// handlerSignin handles user authentication and token creation.
func (apiCfg *apiConfig) handlerSignin(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
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

	user, err := apiCfg.DB.GetUserByUsername(r.Context(), params.Username)

	if err != nil {
		responseWithError(w, http.StatusNotFound, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	if user.Password != params.Password {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Wrong password"))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))

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
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	user, err := apiCfg.DB.GetUserByID(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

// handlerGetPostForUser retrieves posts for a specific user.
func (apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
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
