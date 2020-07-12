package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/tmpchat/tmpchat/message_broker/domain"
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

	repo := repository.NewChatRoomRepository(client)

	message := domain.ChatMessage{ID: "message", Value: "Hello!!", CreatedAt: time.Now()}

	// create room
	room, err := repo.Create()
	if err != nil {
		fmt.Println(err)
	}

	// add message
	err = repo.AddMessage(room.ID, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get messages
	getedMessages, err := repo.GetMessages(room.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, message := range getedMessages {
		fmt.Println(message.Value)
	}

	fmt.Println("Done")
}
