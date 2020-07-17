package broker

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/tmpchat/tmpchat/message_broker/domain"
	"github.com/tmpchat/tmpchat/message_broker/repository"
)

// TODO: change local scope
var upgrader = websocket.Upgrader{} // use default options

type ChatMessageBroker struct {
	repo repository.ChatRoomRepository
}

func NewChatMessageBroker(repo repository.ChatRoomRepository) *ChatMessageBroker {
	return &ChatMessageBroker{repo}
}

func (bro ChatMessageBroker) PostMessage(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
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
		room, err := bro.repo.Find(roomID)
		if err != nil {
			room, err = bro.repo.Create(roomID)
			if err != nil {
				log.Println("create failed.")
			}
		}

		chatmessage := domain.ChatMessage{ID: "id", Value: string(message), CreatedAt: time.Now()}
		// add message
		err = bro.repo.AddMessage(room.ID, chatmessage)
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
