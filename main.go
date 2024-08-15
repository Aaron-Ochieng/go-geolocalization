package main

import (
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.ArtistHandler)
	http.ListenAndServe(":9090", nil)
}
