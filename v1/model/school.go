package model

import "time"

type School struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Name      string    `json:"name,omitempty"`
	Website   *string   `json:"website,omitempty"`
	Latitude  float32   `json:"latitude,omitempty"`
	Longitude float32   `json:"longitude,omitempty"`
}
