package main

import (
	"net/http"
)

type HealthStatus struct {
	Health  string `json:"health"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Ok      bool   `json:"ok"`
}

func healthController(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, HealthStatus{
		Health:  "Good",
		Message: "Server is working good",
		Error:   false,
		Ok:      true,
	})
}