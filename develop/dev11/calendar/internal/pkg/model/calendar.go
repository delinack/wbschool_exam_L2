package model

import (
	"errors"
	"time"
)

type Event struct {
	Id     int       `json:"id" db:"id"`
	UserId int       `json:"user_id" db:"user_id"`
	Title  string    `json:"title" db:"title"`
	Date   time.Time `json:"date" db:"date"`
}

type UpdateCalendarInput struct {
	Id    int        `json:"id" db:"id"`
	Title *string    `json:"title" db:"title"`
	Date  *time.Time `json:"date" db:"date"`
}

func (i UpdateCalendarInput) Validate() error {
	if i.Title == nil && i.Date == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
