package handler

import (
	"calendar/internal/pkg/service"
	"log"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/create_event", h.createEvent)
	router.HandleFunc("/update_event", h.updateEvent)
	router.HandleFunc("/delete_event", h.deleteEvent)
	router.HandleFunc("/events_for_day", h.eventsForDay)
	router.HandleFunc("/events_for_week", h.eventsForWeek)
	router.HandleFunc("/events_for_month", h.eventsForMonth)

	handler := h.loggingMiddleware(router)

	return handler
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
