package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
	messageBrokerHost := os.Getenv("MESSAGE_BROKER_HOST")
	if messageBrokerHost == "" {
		messageBrokerHost = "127.0.0.1:8081"
	}
	request, err := json.Marshal(domain.CraeteChatRoomRequest{ID: room.UUID})

	res, err := http.Post("http://"+messageBrokerHost+"/room", "application/json", bytes.NewBuffer(request))
	if err != nil {
		// TODO: check StatusCode?
		return err
	}
	fmt.Printf("[status] %d\n", res.StatusCode)

	return nil
}

func (r messageBrokerRepository) DeleteRoom(roomID string) error {
	messageBrokerHost := os.Getenv("MESSAGE_BROKER_HOST")
	if messageBrokerHost == "" {
		messageBrokerHost = "127.0.0.1:8081"
	}
	req, err := http.NewRequest(http.MethodDelete, "http://"+messageBrokerHost+"/room/"+roomID, nil)
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
