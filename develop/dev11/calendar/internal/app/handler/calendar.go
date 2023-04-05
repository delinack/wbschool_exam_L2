package handler

import (
	"calendar/internal/pkg/model"
	"calendar/internal/pkg/serializer"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	r.ParseForm()

	event.Title = r.PostForm.Get("title")
	date, err := time.Parse("2006-01-02", r.PostForm.Get("date"))
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}
	event.Date = date
	userId, _ := strconv.Atoi(r.PostForm.Get("user_id"))
	event.UserId = userId

	defer r.Body.Close()

	id, err := h.services.Create(event)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	event.Id = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	var event model.UpdateCalendarInput
	r.ParseForm()

	*event.Title = r.PostForm.Get("title")
	date, err := time.Parse("2006-01-02", r.PostForm.Get("date"))
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}
	*event.Date = date
	defer r.Body.Close()

	err = h.services.Update(event.Id, event)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.PostForm.Get("id")
	if idStr == "" {
		serializer.SerializeError(w, errors.New("empty id field"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.services.Delete(id)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		serializer.SerializeError(w, errors.New("empty date field"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	events, err := h.services.GetForDay(date)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		serializer.SerializeError(w, errors.New("empty date field"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	events, err := h.services.GetForWeek(date)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		serializer.SerializeError(w, errors.New("empty date field"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	events, err := h.services.GetForMonth(date)
	if err != nil {
		serializer.SerializeError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
