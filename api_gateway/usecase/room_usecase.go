package usecase

import (
	google_uuid "github.com/google/uuid"
	"github.com/tmpchat/tmpchat/api_gateway/domain"
	"github.com/tmpchat/tmpchat/api_gateway/gateway"
)

type RoomUsecase interface {
	Create(raw domain.CreateRoomRequest) error
	Find(raw domain.RoomEntity) error
	CreateUUID() google_uuid.UUID
}

type roomUsecase struct {
	roomRepo gateway.RoomRepository
	mbRepo   gateway.MessageBrokerRepository
}

func NewRoomUsecase() RoomUsecase {
	return roomUsecase{
		roomRepo: gateway.NewRoomRepository(),
		mbRepo:   gateway.NewMessageBrokerRepository(),
	}
}

// TODO: Response RoomEntity to Client
func (r roomUsecase) Create(raw domain.CreateRoomRequest) error {
	err := r.roomRepo.Create(raw)
	if err != nil {
		return err
	}

	// TODO: Create MessageBroker
	if err := r.mbRepo.CreateRoom(raw); err != nil {
		return err
	}

	return nil
}

func (r roomUsecase) Find(raw domain.RoomEntity) error {
	_, err := r.roomRepo.Find(raw.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (r roomUsecase) CreateUUID() google_uuid.UUID {
	return google_uuid.UUID(google_uuid.New())
}
