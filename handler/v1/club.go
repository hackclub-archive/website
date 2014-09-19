package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/model"
)

// CreateClub creates a new club from provided JSON.
func CreateClub(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil || u.Type != model.UserAdmin {
		return ErrNotAuthorized()
	}

	defer r.Body.Close()
	club, err := model.NewClub(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	// TODO: Do this in model.NewClub. With the current architecture we have to
	// do this here because using the database package in the model package
	// causes an import cycle.
	_, err = database.GetSchool(club.SchoolID)
	if err == sql.ErrNoRows {
		return ErrCreatingModel(model.ErrInvalidClubSchoolID)
	}

	err = database.SaveClub(club, u)
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, club, http.StatusOK)
}

// GetAllClubs gets all of the clubs from the database. Only administers can
// use this call.
func GetAllClubs(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil || u.Type != model.UserAdmin {
		return ErrNotAuthorized()
	}

	users, err := database.GetClubs()
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, users, http.StatusOK)
}

// GetClub gets a club specified by id.
func GetClub(w http.ResponseWriter, r *http.Request, u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID(err)
	}

	var club *model.Club
	if u.Type == model.UserAdmin {
		club, err = database.GetClub(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrNotFound(err)
			}
			return ErrDatabase(err)
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
func CreateClubMember(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	if u.Type == model.UserStudent {
		return ErrForbidden()
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID(err)
	}

	var club *model.Club
	if u.Type == model.UserAdmin {
		club, err = database.GetClub(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ErrNotFound(err)
			}
			return ErrDatabase(err)
		}
	} else {
		club, err = database.GetClubForUser(id, u.ID)
		if err != nil {
			return ErrNotAuthorized()
		}
	}

	defer r.Body.Close()
	user, err := model.NewUserGeneratePassword(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	if u.Type == model.UserOrganizer && user.Type == model.UserOrganizer {
		return ErrForbidden()
	}

	// TODO: Send email to user

	err = database.SaveUser(user)
	if err != nil {
		if err == model.ErrInvalidUserEmail {
			return ErrCreatingModel(err)
		}
		return ErrDatabase(err)
	}

	err = database.AddUserToClub(user.ID, club.ID)
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, user, http.StatusOK)
}

// GetAllClubsForUser gets all of the clubs that the given user has an
// association with.
func GetAllClubsForUser(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
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

		if id != u.ID {
			return ErrNotAuthorized()
		}
	}

	users, err := database.GetClubsForUser(id)
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, users, http.StatusOK)
}
