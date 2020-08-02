package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tmpchat/tmpchat/message_broker/broker"
)

func main() {
	hub := broker.NewClientHub()
	go hub.Run()
	broker := broker.NewChatMessageBroker(hub)
	http.HandleFunc("/broker", broker.PostMessage)
	http.HandleFunc("/room", broker.CreateRoom)
	log.Fatal(http.ListenAndServe(":8081", nil))

	fmt.Println("Done")
}
