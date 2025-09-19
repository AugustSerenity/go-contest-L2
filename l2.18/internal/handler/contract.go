package handler

import "l2.18/internal/model"

type Service interface {
	CreateEvent(id int, event model.Event) error
}
