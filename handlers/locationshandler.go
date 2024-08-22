package handlers

import (
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := utils.GetLocations(utils.GetApiIndex().Locations)
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		return
	}
	usernames, err := utils.GetArtists()
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	for i, val := range usernames {
		locations[i].Name = val.Name
		// fmt.Println(val.Name)
	}
	tmpl.Execute(w, locations)
}
