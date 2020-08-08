package usecase

import (
	google_uuid "github.com/google/uuid"
	"github.com/tmpchat/tmpchat/api_gateway/domain"
	"github.com/tmpchat/tmpchat/api_gateway/gateway"
)

type RoomUsecase interface {
	CreateRoom(raw domain.CreateRoomRequest) error
	CreateUUID() google_uuid.UUID
}

type roomUsecase struct {
	roomRepo gateway.RoomRepository
	mbRepo   gateway.MessageBrokerRepository
}

func NewRoomUsecase() RoomUsecase {
	return roomUsecase{roomRepo, mbRepo}
}

func (r roomUsecase) CreateRoom(raw domain.CreateRoomRequest) {
	// TODO: check room exists
	room_exists, err := r.roomRepo.Find(room) //
	// TODO: if room is not exists, create room

	err := r.roomRepo.Create(room)
	if err != nil {
		return err
	}
	// TODO: Response RoomEntity to Client
	fmt.Printf(`RoomController.Create: %#v, %#v`, w, r)
}

func (r roomUsecase) CreateUUID() google_uuid.UUID {
	return google_uuid.UUID(google_uuid.New())
}
