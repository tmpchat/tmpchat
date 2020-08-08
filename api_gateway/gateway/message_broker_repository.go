package gateway

import (
	"github.com/tmpchat/tmpchat/api_gateway/domain"
)

type MessageBrokerRepository interface {
	CreateRoom(room domain.CreateRoomRequest) error
}
