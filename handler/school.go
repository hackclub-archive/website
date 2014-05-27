package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/model"
)

// CreateSchool creates a school from JSON in the request body.
func CreateSchool(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil || u.Type != model.UserAdmin {
		err := errors.New("not authorized")
		return &AppError{err, err.Error(), http.StatusUnauthorized}
	}

	defer r.Body.Close()
	school, err := model.NewSchool(r.Body)
	if err != nil {
		return &AppError{err, err.Error(), http.StatusBadRequest}
	}

	err = database.SaveSchool(school)
	if err != nil {
		return &AppError{err, "error saving to database",
			http.StatusInternalServerError}
	}

	return renderJSON(w, school, http.StatusOK)
}

// GetSchool returns the school with the specified ID.
func GetSchool(w http.ResponseWriter, r *http.Request,
	_ *model.User) *AppError {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return &AppError{err, "invalid id", http.StatusBadRequest}
	}

	school, err := database.GetSchool(id)
	if err == sql.ErrNoRows {
		return &AppError{err, "school not found", http.StatusNotFound}
	} else if err != nil {
		return &AppError{err, "error fetching school",
			http.StatusInternalServerError}
	}

	return renderJSON(w, school, http.StatusOK)
}

// GetSchools returns a list of all of the schools.
func GetSchools(w http.ResponseWriter, r *http.Request,
	_ *model.User) *AppError {
	schools, err := database.GetSchools()
	if err != nil {
		return &AppError{err, "error fetching schools",
			http.StatusInternalServerError}
	}

	return renderJSON(w, schools, http.StatusOK)
}
