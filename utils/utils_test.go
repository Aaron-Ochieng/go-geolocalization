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

func TestGetArtists(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.String() != "/api/artists" {
			t.Errorf("Expected to request '/api/artists', got: %s", r.URL.String())
		}

		// Create mock artists data based on the provided sample
		mockArtists := []utils.Artists{
			{
				Id:           1,
				Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
				Name:         "Queen",
				Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
				CreationDate: 1970,
				FirstAlbum:   "14-12-1973",
				Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
				ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
				Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
			},
			{
				Id:           2,
				Image:        "https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
				Name:         "SOJA",
				Members:      []string{"Jacob Hemphill", "Bob Jefferson", "Ryan \"Byrd\" Berty", "Ken Brownell", "Patrick O'Shea", "Hellman Escorcia", "Rafael Rodriguez", "Trevor Young"},
				CreationDate: 1997,
				FirstAlbum:   "05-06-2002",
				Locations:    "https://groupietrackers.herokuapp.com/api/locations/2",
				ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/2",
				Relations:    "https://groupietrackers.herokuapp.com/api/relation/2",
			},
		}

		// Write the mock data as response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockArtists)
	}))
	defer server.Close()

	// Call the function we're testing
	result, err := utils.GetArtists(server.URL + "/api/artists")
	// Check for errors
	if err != nil {
		t.Fatalf("GetArtists() returned an error: %v", err)
	}

	// Define the expected result
	expected := []utils.Artists{
		{
			Id:           1,
			Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
			Name:         "Queen",
			Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
			CreationDate: 1970,
			FirstAlbum:   "14-12-1973",
			Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
			ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
			Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
		},
		{
			Id:           2,
			Image:        "https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
			Name:         "SOJA",
			Members:      []string{"Jacob Hemphill", "Bob Jefferson", "Ryan \"Byrd\" Berty", "Ken Brownell", "Patrick O'Shea", "Hellman Escorcia", "Rafael Rodriguez", "Trevor Young"},
			CreationDate: 1997,
			FirstAlbum:   "05-06-2002",
			Locations:    "https://groupietrackers.herokuapp.com/api/locations/2",
			ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/2",
			Relations:    "https://groupietrackers.herokuapp.com/api/relation/2",
		},
	}

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetArtists() = %v, want %v", result, expected)
	}
}

// ... TestGetArtistsError function remains the same ...
func TestGetArtistsError(t *testing.T) {
	// Test with invalid URL
	_, err := utils.GetArtists("http://invalid-url")
	if err == nil {
		t.Error("Expected error with invalid URL, got nil")
	}

	// Test with server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err = utils.GetArtists(server.URL)
	if err == nil {
		t.Error("Expected error with server error, got nil")
	}

	// Test with invalid JSON
	serverInvalidJSON := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer serverInvalidJSON.Close()

	_, err = utils.GetArtists(serverInvalidJSON.URL)
	if err == nil {
		t.Error("Expected error with invalid JSON, got nil")
	}
}
