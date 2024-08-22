package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Printf("error: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists)
}
