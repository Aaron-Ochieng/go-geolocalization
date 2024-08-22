package handlers

import (
	"net/http"
	"regexp"
)

func HandleUrls(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == "GET" {
			ArtistHandler(w, r)
		} else {
			ErrorHandler(w, r, 405)
		}
	case "/locations":
		if r.Method == "GET" {
			LocationsHandler(w, r)
		} else {
			ErrorHandler(w, r, 405)
		}
	case "/dates":
		if r.Method == "GET" {
			DatesHandler(w, r)
		} else {
			ErrorHandler(w, r, 405)
		}
	case "/search":
		if r.Method == "GET" {
			SearchHandler(w, r)
		} else {
			ErrorHandler(w, r, 405)
		}
	default:
		matched, _ := regexp.MatchString(`^/details/\d+$`, r.URL.Path)
		if matched {
			if r.Method == "GET" {
				DetailsHandler(w, r)
			} else {
				ErrorHandler(w, r, 405)
			}
		} else {
			ErrorHandler(w, r, 400)
		}
	}
}
