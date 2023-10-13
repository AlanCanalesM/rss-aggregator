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
func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
		ID     uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), params.ID)
	if err != nil {
		responseWithError(w, http.StatusNotFound, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreatedFeedFollow(r.Context(), database.CreatedFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not create this feed follow: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

// handlerGetFeedFollows handles the retrieval of feed follows for a specific user.
func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userID")
	userIdUUID, err := uuid.Parse(userIdStr)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing UUID: %v", err))
		return
	}

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), userIdUUID)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not retrieve feed follows: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

// handlerDeleteFeedFollows handles the deletion of a feed follow entry.
func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowIDUUID, err := uuid.Parse(feedFollowIDStr)

	userIdStr := chi.URLParam(r, "userID")
	userIdUUID, err := uuid.Parse(userIdStr)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse the feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowIDUUID,
		UserID: userIdUUID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not delete the feed follow: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, struct{}{})
}
