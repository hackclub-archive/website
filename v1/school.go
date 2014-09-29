package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"code.google.com/p/go.net/context"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/v1/database"
	"github.com/hackedu/backend/v1/school"
	"github.com/hackedu/backend/v1/user"
)

// CreateSchool creates a school from JSON in the request body.
func CreateSchool(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	u, ok := user.FromContext(ctx)
	if !ok || u.Type != user.Admin {
		return ErrNotAuthorized()
	}

	defer r.Body.Close()
	school, err := school.NewSchool(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	err = database.SaveSchool(school)
	if err != nil {
		return err
	}

	return renderJSON(w, school, http.StatusOK)
}

// GetSchool returns the school with the specified ID.
func GetSchool(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID()
	}

	school, err := database.GetSchool(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound()
		}
		return err
	}

	return renderJSON(w, school, http.StatusOK)
}

// GetSchools returns a list of all of the schools.
func GetSchools(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	schools, err := database.GetSchools()
	if err != nil {
		return err
	}

	return renderJSON(w, schools, http.StatusOK)
}
