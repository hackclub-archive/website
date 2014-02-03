package hackedu

import (
	"bytes"
	"testing"

	"appengine/aetest"
	"appengine/datastore"
)

func TestRegisterValidUser(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	name := "foo"
	email := "foo@bar.com"
	password := "foobarfoobar"

	user, key, err := RegisterUser(c, name, email, password)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	user2 := &User{}
	if err = datastore.Get(c, key, user2); err != nil {
		t.Errorf("Failed to retrieve user from datastore. Err: %s", err.Error())
	}

	if user.Name != user2.Name || user.Email != user2.Email ||
		!bytes.Equal(user.Password, user2.Password) {
		t.Errorf("Retrieved invalid user. Expected: %+v, actual %+v", user, user2)
	}
}
