package usecase

import (
	google_uuid "github.com/google/uuid"
)

type RoomUsecase interface {
	CreateUUID() google_uuid.UUID
}

type roomUsecase struct{}

func NewRoomUsecase() RoomUsecase {
	return roomUsecase{}
}

func (r roomUsecase) CreateUUID() google_uuid.UUID {
	return google_uuid.UUID(google_uuid.New())
}
