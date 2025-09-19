package handler

import (
	"encoding/json"
	"net/http"

	"l2.18/internal/handler/response"
	"l2.18/internal/middleware"
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

	router.HandleFunc("POST /create_event", middleware.LoggingMiddleware(h.CreateEvent))

	return router
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var eventRequest model.Request
	err := json.NewDecoder(r.Body).Decode(&eventRequest)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, "bad request")
		return
	}

	event, err := model.CastToEvent(eventRequest)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.CreateEvent(eventRequest.UserID, event)
	if err != nil {
		response.SendError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	response.SendSuccess(w, http.StatusOK, "successfully created")

}
