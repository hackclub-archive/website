package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"code.google.com/p/go.crypto/bcrypt"
	"code.google.com/p/go.net/context"
	"github.com/gorilla/mux"
	"github.com/hackedu/backend/v1/database"
	"github.com/hackedu/backend/v1/token"
	"github.com/hackedu/backend/v1/user"
	"github.com/hackedu/backend/httputil"
)

// Authenticate checks the provided user information against the information
// in the database. If it all checks out, then a JWT is generated and
// returned.
func Authenticate(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	defer r.Body.Close()

	var requestUser user.RequestUser
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		return err
	}

	userFromDB, err := database.GetUserByEmail(requestUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound()
		}
		return err
	}

	err = userFromDB.ComparePassword(requestUser.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return &httputil.HTTPError{http.StatusBadRequest,
			errors.New("invalid password")}
	} else if err != nil {
		return err
	}

	token, err := token.NewToken(userFromDB)
	if err != nil {
		return err
	}

	return renderJSON(w, token, http.StatusOK)
}

// CreateUser creates a new user from JSON in the request body.
func CreateUser(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	//if u == nil || u.Type != user.UserAdmin {
	//return ErrNotAuthorized()
	//}

	defer r.Body.Close()
	u, err := user.NewUser(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	err = database.SaveUser(u)
	if err != nil {
		if err == user.ErrInvalidUserEmail {
			return ErrCreatingModel(err)
		}
		return err
	}

	return renderJSON(w, u, http.StatusOK)
}

// GetUser gets the user specified by ID in the url. If the user is an admin,
// they can see any profile. If the user is an organizer or a member, they can
// only view their own profile. If they are not authorized, they cannot see
// any profiles.
func GetUser(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok {
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
			return err
		}
	}

	if u.Type == user.Admin {
		return renderJSON(w, u, http.StatusOK)
	} else {
		if id == u.ID {
			return renderJSON(w, u, http.StatusOK)
		}

		return ErrNotAuthorized()
	}
}

// GetCurrentUser gets the current authenticated user.
func GetCurrentUser(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok {
		return ErrNotAuthorized()
	}

	return renderJSON(w, u, http.StatusOK)
}
