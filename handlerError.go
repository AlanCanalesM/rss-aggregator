package main

import (
	"net/http"
)

func handlerError(w http.ResponseWriter, r *http.Request) {

	responseWithJSON(w, 400, "Something went wrong!")
}
