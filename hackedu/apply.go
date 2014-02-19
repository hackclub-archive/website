package hackedu

import (
	"encoding/json"
	"net/http"

	"appengine"
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
	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	if err != nil {
		serveError(c, w, err)
		return
	}

	_, err = RegisterUser(c, &u)
	if err != nil {
		serveError(c, w, err)
		return
	}

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		serveError(c, w, err)
		return
	}
}
