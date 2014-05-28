package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var (
	// ErrInvalidClubName is returned when the club's name is invalid
	ErrInvalidClubName = errors.New("invalid name")
)

// Club represents a club participating in hackEDU.
type Club struct {
	ID       int64     `db:"id"        json:"id"`
	Created  time.Time `db:"created"   json:"created"`
	Updated  time.Time `db:"updated"   json:"updated"`
	SchoolID int64     `db:"school_id" json:"school_id"`
	Name     string    `db:"name"      json:"name"`
}

func NewClub(jsonReader io.Reader) (*Club, error) {
	var club Club
	if err := json.NewDecoder(jsonReader).Decode(&club); err != nil {
		return nil, err
	}

	if err := club.validate(); err != nil {
		return nil, err
	}

	return &club, nil
}

func (c *Club) validate() error {
	switch {
	case len(c.Name) == 0 || len(c.Name) > 255:
		return ErrInvalidClubName
	default:
		return nil
	}
}
