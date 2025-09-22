package storage

import (
	"fmt"
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

func (st *Storage) Create(userID int, event model.Event) {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.events[userID] = append(st.events[userID], event)
}

func (st *Storage) Update(userID int, date time.Time, updated model.Event) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	userEvents, ok := st.events[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}

	for i, e := range userEvents {
		if e.Date.Equal(date) {
			st.events[userID][i] = updated
			return nil
		}
	}

	return fmt.Errorf("event not found")
}

func (st *Storage) ExactEventExists(userID int, date time.Time, name string) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()

	userEvents, ok := st.events[userID]
	if !ok {
		return false
	}

	for _, e := range userEvents {
		if e.Date.Equal(date) && e.Name == name {
			return true
		}
	}

	return false
}

func (st *Storage) EventAtTimeExists(userID int, date time.Time) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()

	userEvents, ok := st.events[userID]
	if !ok {
		return false
	}

	for _, e := range userEvents {
		if e.Date.Equal(date) {
			return true
		}
	}

	return false
}

func (st *Storage) Delete(userID int, date time.Time, name string) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	userEvents, ok := st.events[userID]
	if !ok {
		return fmt.Errorf("user not found")
	}

	for i, e := range userEvents {
		if e.Date.Equal(date) && e.Name == name {
			st.events[userID] = append(userEvents[:i], userEvents[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("event not found")
}
