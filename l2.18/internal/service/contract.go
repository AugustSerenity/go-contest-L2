package service

import (
	"time"

	"l2.18/internal/model"
)

type Storage interface {
	Create(id int, event model.Event)
	EventExists(id int, date time.Time, name string) bool
}
