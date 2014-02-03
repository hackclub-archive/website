package hackedu

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type User struct {
	Name        string
	Email       string
	Password    string
	CreatedAt   time.Time
	SignInCount int
}

func userKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", "user", 0, nil)
}
