package handlers

import (
	"groupie-tracker/utils"
	"net/http"
	"text/template"
)

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := utils.GetLocations(utils.GetApiIndex().Locations)
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, locations)
}
