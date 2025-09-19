package model

import "time"

type Request struct {
	UserID    int       `json:"user_id"`
	EventName string    `json:"event_name"`
	Date      time.Time `json:"date"`
}
