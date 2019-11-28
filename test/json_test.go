package test

import (
	"encoding/json"
	"testing"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func TestJSON(t *testing.T) {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		t.Errorf("JSON marshaling failed: %v", err)
	}
	t.Logf("%s", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}
	t.Logf("%v", titles)
}
