package model

import "time"

type Application struct {
	Id                 int       `json:"id"`
	UserId             int       `json:"userId"`
	CreatedAt          time.Time `json:"created_at"`
	HighSchool         string    `json:"highSchool" binding:"required"`
	InterestingProject string    `json:"interestingProject" binding:"required"`
	SystemHacked       string    `json:"systemHacked" binding:"required"`
	Passion            string    `json:"passion" binding:"required"`
	Story              string    `json:"story" binding:"required"`
	Why                string    `json:"why" binding:"required"`
}
