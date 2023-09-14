package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// This file will contain the functions to respond to the client
// We will have a function to respond with an error
// We will have a function to respond with JSON

func responseWithError(W http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Fatal("Response with 5xx status code")
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(W, code, errResponse{
		Error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal the response %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
