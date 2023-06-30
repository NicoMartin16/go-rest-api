package main

import "net/http"

// HandlerReadiness is a handler for readiness probe
func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something went wrong")
}
