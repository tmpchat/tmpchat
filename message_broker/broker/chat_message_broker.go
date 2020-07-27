package broker

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/tmpchat/tmpchat/message_broker/domain"
	"github.com/tmpchat/tmpchat/message_broker/usecase"
)

// TODO: change local scope
var upgrader = websocket.Upgrader{} // use default options

type ChatMessageBroker struct {
	uscs usecase.RoomUsecase
	hub  *ClientHub
}

func NewChatMessageBroker(hub *ClientHub) *ChatMessageBroker {
	return &ChatMessageBroker{uscs: usecase.NewRoomUsecase(), hub: hub}
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

		roomID := "example_room_id"
		chatmessage := &domain.ChatMessage{ID: uuid.New().String(), Value: string(message), CreatedAt: time.Now()}
		res, err := bro.uscs.AddMessage(roomID, chatmessage)
		if err != nil {
			log.Printf(`FAIL add message: %#v.`, err)
			// TODO: error を client に伝える術がないので、エラーを握り潰してる
			continue
		}

		bro.hub.broadcast <- []byte(res.Value)
	}
}

func (bro ChatMessageBroker) closeClient(client *websocket.Conn) {
	defer client.Close()
	client.WriteMessage(websocket.CloseMessage, []byte{})
	bro.hub.unregister <- client
}
