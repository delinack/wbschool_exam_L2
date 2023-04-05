package storage

import (
	"calendar/internal/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type CalendarStorage struct {
	db *sqlx.DB
}

func NewCalendarStorage(db *sqlx.DB) *CalendarStorage {
	return &CalendarStorage{db: db}
}

func (c *CalendarStorage) Create(event model.Event) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, date) VALUES ($1, $2, $3) RETURNING id", eventTable)
	row := c.db.QueryRow(query, event.UserId, event.Title, event.Date)
	if err := row.Scan(&id); err != nil {
		return 0, err

	}
	return id, nil
}

func (c *CalendarStorage) Update(calendarId int, input model.UpdateCalendarInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *input.Date)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s", eventTable, setQuery)
	args = append(args, calendarId)

	_, err := c.db.Exec(query, args...)
	return err
}

func (c *CalendarStorage) Delete(eventId int) error {
	query := fmt.Sprintf("DELETE FROM %s et WHERE et.id = $1", eventTable)

	_, err := c.db.Exec(query, eventId)

	return err
}

func (c *CalendarStorage) GetForDay(date time.Time) (model.Event, error) {
	var day model.Event

	query := fmt.Sprintf("SELECT * FROM %s et WHERE et.date = $1", eventTable)
	err := c.db.Get(&day, query, date)

	return day, err
}

func (c *CalendarStorage) GetForWeek(date time.Time) ([]model.Event, error) {
	var startOfWeek time.Time

	if date.Weekday() == time.Sunday {
		startOfWeek = date.AddDate(0, 0, -6)
	} else {
		daysUntilMonday := int(time.Monday - date.Weekday())
		startOfWeek = date.AddDate(0, 0, daysUntilMonday)
	}

	var week []model.Event
	err := c.db.Select(&week, "SELECT * FROM $s et WHERE et.date >= $1 AND et.date <= $2 ORDER BY et.date ASC",
		eventTable, startOfWeek, startOfWeek.AddDate(0, 0, 6))
	if err != nil {
		return nil, err
	}

	return week, nil
}

func (c *CalendarStorage) GetForMonth(date time.Time) ([]model.Event, error) {
	var startOfMonth = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	var endOfMonth = startOfMonth.AddDate(0, 1, -1)

	var month []model.Event
	err := c.db.Select(&month, "SELECT * FROM %s et WHERE et.date >= $1 AND et.date <= $2 ORDER BY et.date ASC",
		eventTable, startOfMonth, endOfMonth)
	if err != nil {
		return nil, err
	}

	return month, nil
}
