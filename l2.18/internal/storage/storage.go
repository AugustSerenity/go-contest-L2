package storage

import (
	"sync"

	"l2.18/internal/model"
)

type Storage struct {
	mu     sync.RWMutex
	events map[int]model.Event
}

func New() *Storage {
	return &Storage{
		events: make(map[int]model.Event),
	}
}

func (st *Storage) Create(id int, event model.Event) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.events[id] = event
}
