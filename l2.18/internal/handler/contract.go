package handler

import (
	"time"

	"l2.18/internal/model"
)

type Service interface {
	CreateEvent(id int, event model.Event) error
	UpdateEvent(userID int, date time.Time, updated model.Event) error
}
