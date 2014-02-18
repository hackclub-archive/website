package hackedu

import (
	"encoding/json"
	"fmt"
	"net/http"

	"appengine"
)

type School struct {
	Name     string `json:"latitude,omitempty"`
	Location Location
}

type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

func Schools(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	austin := School{
		Name: "Austin High School",
		Location: Location{
			Latitude:  30.27382,
			Longitude: -97.76745,
		},
	}

	thunderridge := School{
		Name: "Thunderridge High School",
		Location: Location{
			Latitude:  39.5347968,
			Longitude: -105.01200670000003,
		},
	}

	schools := []School{austin, thunderridge}

	bytes, err := json.Marshal(schools)
	if err != nil {
		serveError(c, w, err)
	}

	fmt.Println(string(bytes))

	w.Write(bytes)
}
