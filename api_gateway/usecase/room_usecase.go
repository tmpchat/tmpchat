package usecase

import (
	"github.com/tmpchat/tmpchat/api_gateway/domain"
	"github.com/tmpchat/tmpchat/api_gateway/gateway"
)

type RoomUsecase interface {
	Create(req domain.CreateRoomRequest) (*domain.RoomEntity, error)
	Find(id string) (*domain.RoomEntity, error)
	List() ([]*domain.RoomEntity, error)
	UpdateTitle(req domain.UpdateTitleRequest) (*domain.RoomEntity, error)
	Delete(req domain.DeleteRoomRequest) error
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
func (r roomUsecase) Create(req domain.CreateRoomRequest) (*domain.RoomEntity, error) {
	room, err := r.roomRepo.Create(req)
	if err != nil {
		return nil, err
	}

	// TODO: Deactivate MySQL record when failed call MessageBroker API.
	if err := r.mbRepo.CreateRoom(req); err != nil {
		return nil, err
	}

	return room, nil
}

func (r roomUsecase) Find(id string) (*domain.RoomEntity, error) {
	room, err := r.roomRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r roomUsecase) List() ([]*domain.RoomEntity, error) {
	rooms, err := r.roomRepo.List()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r roomUsecase) UpdateTitle(req domain.UpdateTitleRequest) (*domain.RoomEntity, error) {
	room, err := r.roomRepo.UpdateTitle(req)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r roomUsecase) Delete(req domain.DeleteRoomRequest) error {
	err := r.roomRepo.Delete(req)
	if err != nil {
		return err
	}

	return nil
}
