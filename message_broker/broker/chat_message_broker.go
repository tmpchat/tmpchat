package broker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/tmpchat/tmpchat/message_broker/repository"
	"github.com/tmpchat/tmpchat/message_broker/domain"
)

// TODO: change local scope
var upgrader = websocket.Upgrader{} // use default options

type ChatMessageBroker struct {
	repo repository.ChatRoomRepository
}

func NewChatMessageBroker(repo repository.ChatRoomRepository) *ChatMessageBroker {
	return &ChatMessageBroker{repo}
}

func (c ChatMessageBroker) PostMessage(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		// TODO: create interctor? usecase? layer
		roomID := "example_room_id"
		room, err := repo.Find(roomID)
		if err != nil {
			room, err = repo.Create(roomID)
			if err != nil {
				log.Println("create failed.")
			}
		}

		message := domain.ChatMessage{ID: "id", Value: message, CreatedAt: time.Now()}
		// add message
		err = repo.AddMessage(room.ID, message)
		if err != nil {
			log.Println("add failed.")
			break
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

