package broker

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type ClientHub struct {
	// Registered clients.
	clients map[*websocket.Conn]bool

	// Inbound messages from the client.
	broadcast chan []byte

	// Register requests from the client.
	register chan *websocket.Conn

	// Unregister requests from client.
	unregister chan *websocket.Conn
}

func NewClientHub() *ClientHub {
	return &ClientHub{
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]bool),
	}
}


func (h *ClientHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				// close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				// select {
				// case client.send <- message:
				// default:
				// 	close(client.send)
				// 	delete(h.clients, client)
				// }
				// TODO: message type
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					fmt.Println("write message error")
				}
			}
		}
	}
}
