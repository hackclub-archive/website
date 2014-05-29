package handler

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
