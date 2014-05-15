package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
)

func School(w http.ResponseWriter, r *http.Request) *AppError {
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
