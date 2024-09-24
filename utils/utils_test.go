package utils_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"groupie-tracker/utils"
)

var url = "https://groupietrackers.herokuapp.com/api"

func TestGetApiIndex(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.String() != "/api" {
			t.Errorf("Expected to request '/api', got: %s", r.URL.String())
		}

		// Create a mock ApiIndex
		mockApiIndex := utils.ApiIndex{
			Artists:   "https://groupietrackers.herokuapp.com/api/artists",
			Locations: "https://groupietrackers.herokuapp.com/api/locations",
			Dates:     "https://groupietrackers.herokuapp.com/api/dates",
			Relation:  "https://groupietrackers.herokuapp.com/api/relation",
		}

		// Write the mock data as response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockApiIndex)
	}))
	defer server.Close()

	// Save the original URL and replace it with our test server URL
	originalURL := url
	url = server.URL + "/api"
	defer func() { url = originalURL }()

	// Call the function we're testing
	result := utils.GetApiIndex()

	// Define the expected result
	expected := utils.ApiIndex{
		Artists:   "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Dates:     "https://groupietrackers.herokuapp.com/api/dates",
		Relation:  "https://groupietrackers.herokuapp.com/api/relation",
	}

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetApiIndex() = %v, want %v", result, expected)
	}
}
