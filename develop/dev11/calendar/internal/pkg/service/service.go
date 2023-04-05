package service

import (
	"calendar/internal/pkg/model"
	"calendar/internal/pkg/storage"
	"time"
)

type Calendar interface {
	Create(event model.Event) (int, error)
	Update(calendarId int, input model.UpdateCalendarInput) error
	Delete(calendarId int) error
	GetForDay(date time.Time) (model.Event, error)
	GetForWeek(date time.Time) ([]model.Event, error)
	GetForMonth(date time.Time) ([]model.Event, error)
}

type Service struct {
	Calendar
}

func NewService(storage *storage.CalendarStorage) *Service {
	return &Service{
		Calendar: NewCalendarService(storage),
	}
}
