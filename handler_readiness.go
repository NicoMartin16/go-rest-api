package main

import "net/http"

// HandlerReadiness is a handler for readiness probe
func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}
