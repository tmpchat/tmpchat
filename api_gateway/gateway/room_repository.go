package gateway

import (
	"github.com/tmpchat/tmpchat/api_gateway/domain"
)

type RoomRepository interface {
	Create(room domain.RoomEntity) error
	Find(id string) ([]domain.RoomEntity, error)
	List() ([]domain.RoomEntity, error)
	UpdateTitle(id, title string) (domain.RoomEntity, error)
}
