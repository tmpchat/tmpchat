package broker

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/tmpchat/tmpchat/message_broker/domain"
	"github.com/tmpchat/tmpchat/message_broker/usecase"
)

// TODO: change local scope
var upgrader = websocket.Upgrader{} // use default options

type ChatMessageBroker struct {
	uscs usecase.ChatRoomUsecase
	hub  *ClientHub
}

func NewChatMessageBroker(hub *ClientHub) *ChatMessageBroker {
	return &ChatMessageBroker{uscs: usecase.NewChatRoomUsecase(), hub: hub}
}

func (bro ChatMessageBroker) CreateRoom(w http.ResponseWriter, r *http.Request) {
	// TODO: http
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		log.Print("read body error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req, err := domain.DecodeCreateChatRoomRequest(body)
	if err != nil {
		log.Print("decode request error")
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bro.uscs.CreateRoom(req.ID)
	// TODO: http response
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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

func (bro ChatMessageBroker) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, exists := vars["id"]; !exists {
		http.Error(w, "please specify room id", http.StatusBadRequest)
		return
	}
	uuid := vars["id"]

	if err := bro.uscs.DeleteRoom(uuid); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (bro ChatMessageBroker) closeClient(client *websocket.Conn) {
	defer client.Close()
	client.WriteMessage(websocket.CloseMessage, []byte{})
	bro.hub.unregister <- client
}
