package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"code.google.com/p/go.net/context"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/v1/database"
	"github.com/hackedu/backend/v1/club"
	"github.com/hackedu/backend/v1/user"
)

// CreateClub creates a new club from provided JSON.
func CreateClub(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok || u.Type != user.Admin {
		return ErrNotAuthorized()
	}

	defer r.Body.Close()
	c, err := club.NewClub(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	// TODO: Do this in user.NewClub. With the current architecture we have to
	// do this here because using the database package in the user package
	// causes an import cycle.
	//
	// TODO: Take another look at this. Separating all of the users into their
	// own packages may have solved our problem.
	_, err = database.GetSchool(c.SchoolID)
	if err == sql.ErrNoRows {
		return ErrCreatingModel(club.ErrInvalidSchoolID)
	}

	err = database.SaveClub(c, u)
	if err != nil {
		return err
	}

	return renderJSON(w, c, http.StatusOK)
}

// GetAllClubs gets all of the clubs from the database. Only administers can
// use this call.
func GetAllClubs(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok || u.Type != user.Admin {
		return ErrNotAuthorized()
	}

	users, err := database.GetClubs()
	if err != nil {
		return err
	}

	return renderJSON(w, users, http.StatusOK)
}

// GetClub gets a club specified by id.
func GetClub(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID()
	}

	var club *club.Club
	if u.Type == user.Admin {
		club, err = database.GetClub(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrNotFound()
			}
			return err
		}
	} else {
		club, err = database.GetClubForUser(id, u.ID)
		if err != nil {
			return ErrNotAuthorized()
		}
	}

	return renderJSON(w, club, http.StatusOK)
}

// CreateClubMember creates a new user account of type student and adds them
// to the club specified by the URL.
//
// The user's password is generated and emailed to them.
func CreateClubMember(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok {
		return ErrNotAuthorized()
	}

	if u.Type == user.Student {
		return ErrForbidden()
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID()
	}

	var c *club.Club
	if u.Type == user.Admin {
		c, err = database.GetClub(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrNotFound()
			}
			return err
		}
	} else {
		c, err = database.GetClubForUser(id, u.ID)
		if err != nil {
			return ErrNotAuthorized()
		}
	}

	defer r.Body.Close()
	newUser, err := user.NewUserGeneratePassword(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	if u.Type == user.Organizer && newUser.Type == user.Organizer {
		return ErrForbidden()
	}

	// TODO: Send email to user

	err = database.SaveUser(newUser)
	if err != nil {
		if err == user.ErrInvalidUserEmail {
			return ErrCreatingModel(err)
		}
		return err
	}

	err = database.AddUserToClub(newUser.ID, c.ID)
	if err != nil {
		return err
	}

	return renderJSON(w, u, http.StatusOK)
}

// GetAllClubsForUser gets all of the clubs that the given user has an
// association with.
func GetAllClubsForUser(ctx context.Context, w http.ResponseWriter,
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
			return ErrInvalidID()
		}

		if id != u.ID {
			return ErrNotAuthorized()
		}
	}

	users, err := database.GetClubsForUser(id)
	if err != nil {
		return err
	}

	return renderJSON(w, users, http.StatusOK)
}
