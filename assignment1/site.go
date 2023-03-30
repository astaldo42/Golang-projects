package main

import (
	"encoding/json"
	"net/http"
)

func storeHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Serve the store
	store := Store(Items)
	json.NewEncoder(w).Encode(store)
}
