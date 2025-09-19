package service

import (
	"fmt"
	"time"

	"l2.18/internal/model"
)

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{
		storage: st,
	}
}

func (s *Service) CreateEvent(id int, event model.Event) error {
	if !event.Date.After(time.Now()) {
		return fmt.Errorf("past date")
	}
	s.storage.Create(id, event)

	return nil
}
