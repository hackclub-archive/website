package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/model"
)

// AppError represents an error as returned by this application. It works in
// tandem with AppHandler for easy handling of errors.
type AppError struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// AppHandler is a type that implements http.Handler and makes handling
// errors easier. When its method returns an error, it prints it to the logs
// and shows a JSON formatted error to the user.
type AppHandler func(http.ResponseWriter, *http.Request, *model.User) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var e *AppError

	if r.Header.Get("Authorization") != "" {
		var user *model.User
		user, e = getUserFromToken(r)
		if e == nil {
			e = fn(w, r, user)
		}
	} else {
		e = fn(w, r, nil)
	}

	if e != nil { // e is *AppError, not os.Error
		log.Println(e.Error)
		renderJSON(w, e, e.Code)
	}
}

func getUserFromToken(r *http.Request) (*model.User, *AppError) {
	token, err := jwt.ParseFromRequest(r, func(t *jwt.Token) ([]byte, error) {
		// TODO: Use real secret
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, &AppError{err, "bad authorization token",
			http.StatusBadRequest}
	}

	userID := int64(token.Claims["id"].(float64))

	user, err := database.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &AppError{err, "user from token not found",
				http.StatusNotFound}
		}
		return nil, &AppError{err, "error fetching user from database",
			http.StatusInternalServerError}
	}

	return user, nil
}

func renderJSON(w http.ResponseWriter, data interface{}, code int) *AppError {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return &AppError{err, "error encoding json",
			http.StatusInternalServerError}
	}
	return nil
}
