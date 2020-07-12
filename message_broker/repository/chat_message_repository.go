package repository

import (
	"bytes"

	"github.com/go-redis/redis"

	"github.com/tmpchat/tmpchat/message_broker/domain"
)

type ChatMessageRepository interface {
	Get(roomID string) (*domain.ChatMessage, error)
	Set(roomID string, message domain.ChatMessage) error
}

type chatMessageRepository struct {
	db *redis.Client
}

// NewChatMessageRepository create ChatMessageRepository
func NewChatMessageRepository(client *redis.Client) ChatMessageRepository {
	return chatMessageRepository{client}
}

func (c chatMessageRepository) Get(roomID string) (*domain.ChatMessage, error) {
	raw, err := c.db.Get(roomID).Result()
	if err != nil {
		return nil, err
	}

	message := domain.ChatMessage{}
	if err := message.UnmarshalBinary(bytes.NewBufferString(raw).Bytes()); err != nil {
		return nil, err
	}
	return &message, nil
}

func (c chatMessageRepository) Set(roomID string, message domain.ChatMessage) error {
	err := c.db.Set(roomID, message, 0).Err()
	return err
}
