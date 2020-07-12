package main

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/tmpchat/tmpchat/message_broker/domain"
)

type ChatMessageRepository interface {
	// TODO: add room id
	Get(id string) (*domain.ChatMessage, error)
	// TODO: add room id
	Set(domain.ChatMessage) (error)
}

type chatMessageRepository struct {
	db *redis.Client
}

// NewChatMessageRepository create ChatMessageRepository
func NewChatMessageRepository(client *redis.Client) ChatMessageRepository {
	return chatMessageRepository{client}
}

func (c chatMessageRepository) Get(id string) (*domain.ChatMessage, error) {
	fmt.Println("not impl Get")
	return nil, nil
}

func (c chatMessageRepository) Set(domain.ChatMessage) (error) {
	fmt.Println("not impl Set")
	return nil
}
