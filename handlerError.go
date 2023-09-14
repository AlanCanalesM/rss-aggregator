package main

import (
	"net/http"
)

// This function will handle the errors
func handlerError(w http.ResponseWriter, r *http.Request) {

	responseWithJSON(w, 400, "Something went wrong!")
}
