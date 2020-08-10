package usecase

import (
	"github.com/tmpchat/tmpchat/api_gateway/domain"
	"github.com/tmpchat/tmpchat/api_gateway/gateway"
)

type RoomUsecase interface {
	Create(raw domain.CreateRoomRequest) error
	Find(id string) (*domain.RoomEntity, error)
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

func (r roomUsecase) Find(id string) (*domain.RoomEntity, error) {
	row, err := r.roomRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return row, err
}
