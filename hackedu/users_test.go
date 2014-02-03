package hackedu

import (
	"testing"

	"appengine/aetest"
	"appengine/datastore"
)

func TestUserKey(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	key := userKey(c)

	if !key.Equal(datastore.NewKey(c, "User", "user", 0, nil)) {
		t.Errorf("Created user key does not match expected.")
	}
}
