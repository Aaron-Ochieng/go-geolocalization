package handlers

import (
	"groupie-tracker/utils"
	"net/http"
	"text/template"
)

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := utils.GetDates(utils.GetApiIndex().Dates)
	if err != nil {
		http.Error(w, "Error fetching dates", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dates)
}
