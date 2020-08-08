package gateway

import "github.com/tmpchat/tmpchat/api_gateway/domain"

type MessageBrokerRepository interface {
	CreateRoom(room domain.CreateRoomRequest) error
}

type messageBrokerRepository struct {
}

func NewMessageBrokerRepository() MessageBrokerRepository {
	return messageBrokerRepository{}
}

func (r messageBrokerRepository) CreateRoom(room domain.CreateRoomRequest) error {
	panic("not impl")
}
