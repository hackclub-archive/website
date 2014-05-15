package model

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"
)

var (
	// ErrInvalidName is returned when the school's name is invalid
	ErrInvalidName = errors.New("invalid name")
	// ErrInvalidWebsite is returned when the school's website is invalid
	ErrInvalidWebsite = errors.New("invalid website")
	// ErrInvalidLatitude is returned when the school's latitude is invalid
	ErrInvalidLatitude = errors.New("invalid latitude")
	// ErrInvalidLongitude is returned when the school's longitude is invalid
	ErrInvalidLongitude = errors.New("invalid longitude")
)

var regexpURL = regexp.MustCompile(`/((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)/`)

// School represents a school that hackEDU is in.
type School struct {
	ID        int64     `db:"id" 		   json:"id"`
	Created   time.Time `db:"created"  json:"created"`
	Updated   time.Time `db:"updated"  json:"updated"`
	Name      string    `db:"name"     json:"name"`
	Website   string    `db:"website"  json:"website"`
	Latitude  float64   `db:"latitude" json:"latitude"`
	Longitude float64   `db:"latitude" json:"latitude"`
}

// NewSchool creates a new school from provided JSON data. It unmarshales the
// JSON, validates the fields, then returns the created school.
//
// NewSchool does not commit the created school to the database.
func NewSchool(jsonData []byte) (*School, error) {
	var school School
	err := json.Unmarshal(jsonData, &school)
	if err != nil {
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
		return ErrInvalidName
	case len(s.Name) > 255:
		return ErrInvalidName
	case regexpURL.MatchString(s.Website) == false:
		return ErrInvalidWebsite
	case s.Latitude == 0:
		return ErrInvalidLatitude
	case s.Longitude == 0:
		return ErrInvalidLongitude
	default:
		return nil
	}
}
