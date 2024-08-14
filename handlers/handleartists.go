package handlers

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/utils"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Replace with your actual URL or get it from a config/environment variable
	url := "http://example.com/api/artists"

	artists := utils.GetArtists(url)
	if artists == nil {
		http.Error(w, "Unable to retrieve artists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(artists); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
