package broker

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type ClientHub struct {
	// Registered clients.
	clients map[string]map[*Client]bool

	// Inbound messages from the client.
	broadcast chan Message

	// Register requests from the client.
	register chan *Client

	// Unregister requests from client.
	unregister chan *Client
}

type Client struct {
	Conn *websocket.Conn
	RoomID string
}

type Message struct {
	Value []byte
	RoomID string
}

func NewClientHub() *ClientHub {
	return &ClientHub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
	}
}


func (h *ClientHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.RoomID][client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client.RoomID][client]; ok {
				delete(h.clients[client.RoomID], client)
				// close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients[message.RoomID] {
				// select {
				// case client.send <- message:
				// default:
				// 	close(client.send)
				// 	delete(h.clients, client)
				// }
				// TODO: message type
				err := client.Conn.WriteMessage(websocket.TextMessage, message.Value)
				if err != nil {
					fmt.Println("write message error")
				}
			}
		}
	}
}
