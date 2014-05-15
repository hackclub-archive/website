package database

import (
	"time"

	"github.com/hackedu/backend/model"
)

// GetSchool gets a school from the database with the provided ID.
func GetSchool(id int64) (*model.School, error) {
	var school model.School
	err := db.Get(&school, "SELECT * FROM schools WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &school, nil
}

// SaveSchool saves the provided school to the database. If the school is a
// new school, then the school.Created field is set to the current time. The
// school.Updated field is set to the current time regardless.
func SaveSchool(school *model.School) error {
	if school.ID == 0 {
		school.Created = time.Now()
	}
	school.Updated = time.Now()

	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO schools (created, updated, name, website, latitude, longitude) VALUES (:created, :updated, :name, :website, :latitude, :longitude)", school)
	tx.Commit()

	return nil
}
