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

// handlerCreateFeedFollows handles the creation of a new feed follow entry.
func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID   uuid.UUID `json:"feed_id"`
		FeedName string    `json:"feed_name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreatedFeedFollow(r.Context(), database.CreatedFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		FeedName:  params.FeedName,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not create this feed follow: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

// handlerGetFeedFollows handles the retrieval of feed follows for a specific user.
func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not retrieve feed follows: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

// handlerDeleteFeedFollows handles the deletion of a feed follow entry.
func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowIDUUID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse the feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feedFollowIDUUID,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not delete the feed follow: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, struct{}{})
}
