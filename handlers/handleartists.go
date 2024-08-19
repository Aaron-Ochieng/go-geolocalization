package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := utils.GetArtists()
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Printf("error: %v", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists)
}
