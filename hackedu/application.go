package hackedu

import (
	"errors"
	"time"

	"appengine"
	"appengine/datastore"
)

type Application struct {
	CreatedAt          time.Time `json:"created_at,omitempty"`
	HighSchool         string    `json:"high_school,omitempty"`
	InterestingProject string    `json:"interesting_project,omitempty"`
	SystemHacked       string    `json:"system_hacked,omitempty"`
	Passion            string    `json:"passion,omitempty"`
	Story              string    `json:"story,omitempty"`
	Why                string    `json:"why,omitempty"`
}

func ConstructApplication(c appengine.Context,
	application *Application) (key *datastore.Key, err error) {

	if !(len(application.HighSchool) > 0) {
		return nil, errors.New("A high school is required.")
	}

	if !(len(application.InterestingProject) > 0) {
		return nil, errors.New("An interesting project is required.")
	}

	if !(len(application.SystemHacked) > 0) {
		return nil, errors.New("A system hacked is required.")
	}

	if !(len(application.Passion) > 0) {
		return nil, errors.New("Your passion is required.")
	}

	if !(len(application.Story) > 0) {
		return nil, errors.New("Your story is required.")
	}

	if !(len(application.Why) > 0) {
		return nil, errors.New("You must say why you want to start a Hack Club.")
	}

	application.CreatedAt = time.Now()

	key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "application",
		nil), application)
	if err != nil {
		return key, err
	}

	return key, nil
}
