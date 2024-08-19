package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.ArtistHandler)
	// The server runs asynchronously on port 8080 using a goroutine.
	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("Server running on http://localhost:8000")
	select {} // keep the main function running indefinitely.
}
