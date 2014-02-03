package hackedu

import (
	"time"

	"code.google.com/p/go.crypto/scrypt"

	"appengine"
	"appengine/datastore"
)

type User struct {
	CreatedAt time.Time
	Name      string
	Email     string
	Password  []byte
}

func RegisterUser(c appengine.Context, name, email, password string) (*User,
	*datastore.Key, error) {
	hashedPassword, err := scrypt.Key([]byte(password), []byte("saltgoeshere"),
		16384, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}

	user := &User{
		CreatedAt: time.Now(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
	}

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "user", nil),
		user)
	if err != nil {
		return user, key, err
	}

	return user, key, nil
}
