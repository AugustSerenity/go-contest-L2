package handler

import (
	"encoding/json"
	"net/http"

	"l2.18/internal/model"
)

type Handler struct {
	service Service
}

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Route() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /create_event", h.CreateEvent)

	return router
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusBadRequest)
		return
	}

	var eventRequest model.Request
	err := json.NewDecoder(r.Body).Decode(&eventRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	event, err := model.CastToEvent(eventRequest)
	if err != nil {
		//...
		return
	}

	err = h.service.CreateEvent(eventRequest.UserID, event)
	if err != nil {
		//...
		return
	}

}
