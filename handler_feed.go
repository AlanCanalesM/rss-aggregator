package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerGetFeedsNotFollowed(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiCfg.DB.FeedsNotFollowedByUser(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))

}

// handlerCreateFeed handles the creation of a feed.
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusNotFound, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

// handlerGetFeeds handles the retrieval of feeds.
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
