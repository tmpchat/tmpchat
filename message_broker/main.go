package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/tmpchat/tmpchat/message_broker/domain"
)

func main() {
	fmt.Println("Hello MessageBroker.")

	// TODO: add redis config
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
	})

	repo := NewChatMessageRepository(client)

	message, err := repo.Get("id")
	if err != nil {
		return
	}

	err = repo.Set(domain.ChatMessage{})
	if err != nil {
		return
	}
}
