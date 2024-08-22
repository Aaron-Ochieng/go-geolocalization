package handlers

import (
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := utils.GetDates(utils.GetApiIndex().Dates)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	users, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	for i, val := range users {
		dates[i].Name = val.Name
	}
	tmpl, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dates)
}
