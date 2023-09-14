package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// This file will contain the handlers for the feed follows in the application
func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreatedFeedFollow(r.Context(), database.CreatedFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not create this feed follows: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseFeedFollowToFeedFollow(feed))

}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not create this feed follows: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))

}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIdStr := chi.URLParam(r, "feedFollowID")
	feedFollowIdUUID, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not parse the feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{

		ID:     feedFollowIdUUID,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not delete the feed follows: %v", err))
		return
	}

	responseWithJSON(w, 200, struct{}{})

}
