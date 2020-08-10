package domain

import (
	"time"

	"github.com/google/uuid"
)

type RoomEntity struct {
	ID        int        `json:"id"`
	UUID      string     `json:"uuid"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateRoomRequest struct {
	UUID  string `json:"-"`
	Title string `json:"title"`
}

func NewCreateRoomRequest() CreateRoomRequest {
	return CreateRoomRequest{UUID: uuid.New().String()}
}
