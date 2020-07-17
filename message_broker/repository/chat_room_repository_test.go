package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/tmpchat/tmpchat/message_broker/domain"
)

func TestBasicUsecase(t *testing.T) {
	// TODO: add redis config
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		t.Error("ping failed.")
	}

	repo := NewChatRoomRepository(client)

	message := domain.ChatMessage{ID: "message", Value: "Hello!!", CreatedAt: time.Now()}

	// create room
	roomID := "example_room_id"
	room, err := repo.Find(roomID)
	if err != nil {
		room, err = repo.Create(roomID)
		if err != nil {
			t.Error("create failed.")
		}
	}

	// add message
	err = repo.AddMessage(room.ID, message)
	if err != nil {
		t.Error("add failed.")
	}

	// get messages
	getedMessages, err := repo.GetMessages(room.ID)
	if err != nil {
		t.Error("get failed.")
	}

	for _, message := range getedMessages {
		fmt.Println(message.Value)
	}

	fmt.Println("Done")
}
