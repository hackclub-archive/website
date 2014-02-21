package hackedu

import (
	"encoding/json"
	"net/http"

	"appengine"
	"appengine/datastore"
)

type ApplicationForm struct {
	Application Application `json:"application,omitempty"`
	User        User        `json:"user,omitempty"`
}

func Apply(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var applicationForm ApplicationForm
	err := decoder.Decode(&applicationForm)
	if err != nil {
		serveError(c, w, err)
		return
	}

	user := applicationForm.User
	application := applicationForm.Application

	userKey, err := RegisterUser(c, &user)
	if err != nil {
		serveError(c, w, err)
		return
	}

	applicationKey, err := ConstructApplication(c, &application)
	if err != nil {
		serveError(c, w, err)
		return
	}

	user.Application = applicationKey

	_, err = datastore.Put(c, userKey, &user)
	if err != nil {
		serveError(c, w, err)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		serveError(c, w, err)
		return
	}
}
