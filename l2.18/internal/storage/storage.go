package storage

import (
	"sync"
	"time"

	"l2.18/internal/model"
)

type Storage struct {
	mu     sync.RWMutex
	events map[int][]model.Event
}

func New() *Storage {
	return &Storage{
		events: make(map[int][]model.Event),
	}
}

func (st *Storage) Create(id int, event model.Event) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.events[id] = append(st.events[id], event)
}

func (st *Storage) EventExists(userID int, date time.Time, name string) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()

	userEvents, ok := st.events[userID]
	if !ok {
		return false
	}

	for _, e := range userEvents {
		if e.Name == name && e.Date.Equal(date) {
			return true
		}
	}

	return false
}
