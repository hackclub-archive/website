package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/model"
)

// School returns the school with the specified ID.
func School(w http.ResponseWriter, r *http.Request, _ *model.User) *AppError {
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

// Schools returns a list of all of the schools.
func Schools(w http.ResponseWriter, r *http.Request, _ *model.User) *AppError {
	schools, err := database.GetSchools()
	if err != nil {
		return &AppError{err, "error fetching schools",
			http.StatusInternalServerError}
	}

	return renderJSON(w, schools, http.StatusOK)
}
