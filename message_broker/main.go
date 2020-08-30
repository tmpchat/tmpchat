package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tmpchat/tmpchat/message_broker/broker"
)

func main() {
	hub := broker.NewClientHub()
	go hub.Run()
	router := mux.NewRouter()
	broker := broker.NewChatMessageBroker(hub)
	router.HandleFunc("/room", broker.CreateRoom).Methods("POST")
	router.HandleFunc("/room/{id}", broker.DeleteRoom).Methods("DELETE")
	router.HandleFunc("/broker/{id}", broker.PostMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

	fmt.Println("Done")
}
