package hackedu

import (
	"bytes"
	"strings"
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

func TestRegisterInvalidUser(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	name := "foo"
	email := "foo@bar.com"
	password := "foobarfoobar"

	_, _, err = RegisterUser(c, "", email, password)
	if !strings.Contains(err.Error(), "name") {
		t.Errorf("Expected error to contain name, actual: %s", err.Error())
	}

	_, _, err = RegisterUser(c, name, "foo", password)
	if !strings.Contains(err.Error(), "email") {
		t.Errorf("Expected error to contain email, actual: %s", err.Error())
	}

	_, _, err = RegisterUser(c, name, email, "short")
	if !strings.Contains(err.Error(), "password") {
		t.Errorf("Expected error to contain password, actual: %s", err.Error())
	}
}
