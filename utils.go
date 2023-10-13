package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) GetUser(w http.ResponseWriter, r *http.Request) database.User {
	type parameters struct {
		ID string `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return database.User{}
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), uuid.MustParse(params.ID))

	if err != nil {
		responseWithError(w, 404, fmt.Sprintf("Could not get user: %v", err))
		return database.User{}
	}

	return user
}
