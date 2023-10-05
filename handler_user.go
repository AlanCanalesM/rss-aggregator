package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// This file will contain the handlers for the user in the application
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	responseWithJSON(w, 200, databaseUserToUser(user))

}

func (apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{

		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		responseWithError(w, 400, "We couldn't reach the posts")
	}

	responseWithJSON(w, 200, databasePostsToPosts(posts))

}

func (apiCfg *apiConfig) handlerGetAllPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := apiCfg.DB.GetAllPosts(r.Context())

	if err != nil {
		responseWithError(w, 400, "We couldn't reach the posts")
	}

	responseWithJSON(w, 200, databasePostsToPosts(posts))

}
