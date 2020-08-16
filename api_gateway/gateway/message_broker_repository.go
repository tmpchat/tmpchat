package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tmpchat/tmpchat/api_gateway/domain"
)

type MessageBrokerRepository interface {
	CreateRoom(room domain.CreateRoomRequest) error
	DeleteRoom(roomID string) error
}

type messageBrokerRepository struct {
}

func NewMessageBrokerRepository() MessageBrokerRepository {
	return messageBrokerRepository{}
}

func (r messageBrokerRepository) CreateRoom(room domain.CreateRoomRequest) error {
	request, err := json.Marshal(domain.CraeteChatRoomRequest{ID: room.UUID})

	res, err := http.Post("http://localhost:8081/room", "application/json", bytes.NewBuffer(request))
	if err != nil {
		// TODO: check StatusCode?
		return err
	}
	fmt.Printf("[status] %d\n", res.StatusCode)

	return nil
}

func (r messageBrokerRepository) DeleteRoom(roomID string) error {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8081/room/"+roomID, nil)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		// TODO: check StatusCode?
		return err
	}
	fmt.Printf("[status] %d\n", res.StatusCode)

	return nil
}
