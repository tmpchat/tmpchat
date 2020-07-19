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
	hub  *ClientHub
}

func NewChatMessageBroker(repo repository.ChatRoomRepository, hub *ClientHub) *ChatMessageBroker {
	return &ChatMessageBroker{repo: repo, hub: hub}
}

func (bro ChatMessageBroker) PostMessage(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Print("receive message")
	bro.hub.register <- c
	defer bro.closeClient(c)
	for {
		_, message, err := c.ReadMessage()
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

		bro.hub.broadcast <- message
	}
}

func (bro ChatMessageBroker) closeClient(client *websocket.Conn) {
	defer client.Close()
	client.WriteMessage(websocket.CloseMessage, []byte{})
	bro.hub.unregister <- client
}
