package model

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"
)

var (
	// ErrInvalidSchoolName is returned when the school's name is invalid
	ErrInvalidSchoolName = errors.New("invalid name")
	// ErrInvalidSchoolWebsite is returned when the school's website is invalid
	ErrInvalidSchoolWebsite = errors.New("invalid website")
	// ErrInvalidSchoolLatitude is returned when the school's latitude is invalid
	ErrInvalidSchoolLatitude = errors.New("invalid latitude")
	// ErrInvalidSchoolLongitude is returned when the school's longitude is
	// invalid
	ErrInvalidSchoolLongitude = errors.New("invalid longitude")
)

var regexpURL = regexp.MustCompile(`/((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)/`)

// School represents a school that hackEDU is in.
type School struct {
	ID        int64     `db:"id"        json:"id"`
	Created   time.Time `db:"created"   json:"created"`
	Updated   time.Time `db:"updated"   json:"updated"`
	Name      string    `db:"name"      json:"name"`
	Website   string    `db:"website"   json:"website"`
	Latitude  float64   `db:"latitude"  json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
}

// NewSchool creates a new school from provided JSON reader. It decodes the
// JSON, validates the fields, then returns the created school.
//
// NewSchool does not commit the created school to the database.
func NewSchool(jsonReader io.Reader) (*School, error) {
	var school School
	if err := json.NewDecoder(jsonReader).Decode(&school); err != nil {
		return nil, err
	}

	if err := school.validate(); err != nil {
		return nil, err
	}

	return &school, nil
}

func (s *School) validate() error {
	switch {
	case len(s.Name) < 6:
		return ErrInvalidSchoolName
	case len(s.Name) > 255:
		return ErrInvalidSchoolName
	case regexpURL.MatchString(s.Website) == false:
		return ErrInvalidSchoolWebsite
	case s.Latitude == 0:
		return ErrInvalidSchoolLatitude
	case s.Longitude == 0:
		return ErrInvalidSchoolLongitude
	default:
		return nil
	}
}
