package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/tmpchat/tmpchat/message_broker/broker"
	"github.com/tmpchat/tmpchat/message_broker/repository"
)

func main() {
	fmt.Println("Hello MessageBroker.")

	// TODO: add redis config
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		fmt.Println("redis connection error")
		return
	}

	repo := repository.NewChatRoomRepository(client)
	broker := broker.NewChatMessageBroker(repo)
	http.HandleFunc("/broker", broker.PostMessage)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))

	fmt.Println("Done")
}
