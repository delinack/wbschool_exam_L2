package service

import (
	"calendar/internal/pkg/model"
	"calendar/internal/pkg/storage"
	"time"
)

type CalendarService struct {
	storage storage.CalendarStorage
}

func NewCalendarService(storage *storage.CalendarStorage) *CalendarService {
	return &CalendarService{storage: *storage}
}

func (c *CalendarService) Create(event model.Event) (int, error) {
	return c.storage.Create(event)
}

func (c *CalendarService) Update(calendarId int, input model.UpdateCalendarInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return c.storage.Update(calendarId, input)
}

func (c *CalendarService) Delete(calendarId int) error {
	return c.storage.Delete(calendarId)
}

func (c *CalendarService) GetForDay(date time.Time) (model.Event, error) {
	return c.storage.GetForDay(date)
}

func (c *CalendarService) GetForWeek(date time.Time) ([]model.Event, error) {
	return c.storage.GetForWeek(date)
}

func (c *CalendarService) GetForMonth(date time.Time) ([]model.Event, error) {
	return c.storage.GetForMonth(date)
}
