package domain

import (
	"time"

	validator "github.com/go-playground/validator/v10"
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
	Title string `validate:"required" json:"title"`
}

type UpdateTitleRequest struct {
	UUID  string `validate:"required" json:"uuid"`
	Title string `validate:"required" json:"title"`
}

func NewCreateRoomRequest() CreateRoomRequest {
	return CreateRoomRequest{UUID: uuid.New().String()}
}

func (req CreateRoomRequest) Validate() error {
	return validator.New().Struct(req)
}

func (req UpdateTitleRequest) Validate() error {
	return validator.New().Struct(req)
}
