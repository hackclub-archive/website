package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/model"
)

// Authenticate checks the provided user information against the information
// in the database. If it all checks out, then a JWT is generated and
// returned.
func Authenticate(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	defer r.Body.Close()

	var requestUser model.RequestUser
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		return ErrUnmarshalling(err)
	}

	userFromDB, err := database.GetUserByEmail(requestUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound(err)
		}
		return ErrDatabase(err)
	}

	err = userFromDB.ComparePassword(requestUser.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return &AppError{err, "invalid password", http.StatusBadRequest}
	} else if err != nil {
		return &AppError{err, "error checking password",
			http.StatusInternalServerError}
	}

	token, err := model.NewToken(userFromDB)
	if err != nil {
		return &AppError{err, "problem creating jwt token",
			http.StatusInternalServerError}
	}

	return renderJSON(w, token, http.StatusOK)
}

// CreateUser creates a new user from JSON in the request body.
func CreateUser(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	//if u == nil || u.Type != model.UserAdmin {
	//return ErrNotAuthorized()
	//}

	defer r.Body.Close()
	user, err := model.NewUser(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	err = database.SaveUser(user)
	if err != nil {
		if err == model.ErrInvalidUserEmail {
			return ErrCreatingModel(err)
		}
		return ErrDatabase(err)
	}

	return renderJSON(w, user, http.StatusOK)
}

// GetUser gets the user specified by ID in the url. If the user is an admin,
// they can see any profile. If the user is an organizer or a member, they can
// only view their own profile. If they are not authorized, they cannot see
// any profiles.
func GetUser(w http.ResponseWriter, r *http.Request, u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	stringID := vars["id"]

	var id int64
	if stringID == "me" {
		id = u.ID
	} else {
		var err error
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			return ErrInvalidID(err)
		}
	}

	if u.Type == model.UserAdmin {
		return renderJSON(w, u, http.StatusOK)
	} else {
		if id == u.ID {
			return renderJSON(w, u, http.StatusOK)
		}

		return ErrNotAuthorized()
	}
}

// GetCurrentUser gets the current authenticated user.
func GetCurrentUser(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	return renderJSON(w, u, http.StatusOK)
}
