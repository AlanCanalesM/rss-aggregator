package main

import (
	"net/http"
)

func handlerHealtCheck(w http.ResponseWriter, r *http.Request) {

	responseWithJSON(w, 200, struct{}{})
}
