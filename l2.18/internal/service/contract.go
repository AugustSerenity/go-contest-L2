package service

import "l2.18/internal/model"

type Storage interface {
	Create(id int, event model.Event)
}
