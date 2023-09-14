package main

import (
	"net/http"
)

// This function will handle the health check endpoint
func handlerHealtCheck(w http.ResponseWriter, r *http.Request) {

	responseWithJSON(w, 200, struct{}{})
}
