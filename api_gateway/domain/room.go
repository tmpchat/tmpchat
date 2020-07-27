package domain

import (
	"time"
)

type RoomEntity struct {
	ID        string
	UUID      string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
