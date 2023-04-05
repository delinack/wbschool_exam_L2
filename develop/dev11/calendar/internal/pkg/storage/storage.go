package storage

import (
	"calendar/internal/pkg/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type Calendar interface {
	Create(event model.Event) (int, error)
	Update(calendarId int, input model.UpdateCalendarInput) error
	Delete(eventId int) error
	GetForDay(date time.Time) (model.Event, error)
	GetForWeek(date time.Time) ([]model.Event, error)
	GetForMonth(date time.Time) ([]model.Event, error)
}

type Storage struct {
	Calendar
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Calendar: NewCalendarStorage(db),
	}
}
