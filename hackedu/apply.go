package hackedu

import (
	"fmt"
	"net/http"
)

type Application struct {
	HighSchool         string `json:"high_school,omitempty"`
	InterestingProject string `json:"interesting_project,omitempty"`
	SystemHacked       string `json:"system_hacked,omitempty"`
	Passion            string `json:"passion,omitempty"`
	Story              string `json:"story,omitempty"`
	Why                string `json:"why,omitempty"`
}

func Apply(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi!")
}
