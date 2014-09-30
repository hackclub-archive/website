package database

import (
	"time"

	"github.com/hackedu/backend/v1/school"
)

const schoolGetByIDStmt = `SELECT id, created, updated, name, website,
latitude, longitude FROM schools WHERE id=$1`

const schoolGetAllStmt = `SELECT id, created, updated, name, website,
latitude, longitude FROM schools ORDER BY id`

const schoolCreateStmt = `INSERT INTO schools (created, updated, name, website, latitude, longitude) VALUES ($1 ,$2, $3, $4, $5, $6) RETURNING id`

// GetSchool gets a school from the database with the provided ID.
func GetSchool(id int64) (*school.School, error) {
	s := new(school.School)
	row := db.QueryRow(schoolGetByIDStmt, id)
	if err := row.Scan(&s.ID, &s.Created, &s.Updated, &s.Name, &s.Website,
		&s.Latitude, &s.Longitude); err != nil {
		return nil, err
	}
	return s, nil
}

// GetSchools gets all of the schools from the database ordered by id.
func GetSchools() ([]*school.School, error) {
	schools := []*school.School{}
	rows, err := db.Query(schoolGetAllStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		s := new(school.School)
		if err := rows.Scan(&s.ID, &s.Created, &s.Updated, &s.Name, &s.Website,
			&s.Latitude, &s.Longitude); err != nil {
			return nil, err
		}

		schools = append(schools, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schools, nil
}

// SaveSchool saves the provided school to the database. If the school is a
// new school, then the school.Created field is set to the current time. The
// school.Updated field is set to the current time regardless.
func SaveSchool(s *school.School) error {
	if s.ID == 0 {
		s.Created = time.Now()
	}
	s.Updated = time.Now()

	row := db.QueryRow(schoolCreateStmt, s.Created, s.Updated, s.Name, s.Website, s.Latitude, s.Longitude)
	if err := row.Scan(&s.ID); err != nil {
		return err
	}

	return nil
}
