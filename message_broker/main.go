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

	repo := repository.NewChatMessageRepository(client)

	_, err = repo.Get("id")
	if err != nil {
		return
	}

	message := domain.ChatMessage{ID: "message",Value:"Hello!!", CreatedAt:time.Now()}
	err = repo.Set("example", message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
