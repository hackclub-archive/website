package route

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"code.google.com/p/go.crypto/scrypt"

	"github.com/coopernurse/gorp"
	"github.com/hackedu/backend/v1/model"
)

func AddUser(user model.User, db gorp.SqlExecutor) (int, string) {
	application := user.Application
	user.Application = nil

	user.CreatedAt = time.Now()
	application.CreatedAt = time.Now()

	salt := user.CreatedAt.String() + user.LastName + user.Email

	hashedPassword, err := hashPassword(user.Password, salt)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, "Error while creating user."
	}

	user.Password = ""
	user.PasswordVerify = ""
	user.HashedPassword = hashedPassword

	// TODO: Figure out why this isn't doing anything.
	db.Insert(&user)

	application.UserId = user.Id

	db.Insert(application)

	// TODO: Send emails

	json, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, "Error while creating user."
	}

	return http.StatusOK, string(json)
}

func hashPassword(password, salt string) ([]byte, error) {
	return scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
}
