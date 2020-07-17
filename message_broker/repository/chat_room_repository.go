package repository

import (
	"bytes"

	"github.com/go-redis/redis"

	"github.com/tmpchat/tmpchat/message_broker/domain"
)

type ChatRoomRepository interface {
	Create() (*domain.ChatRoom, error)
	Find(roomID string) (*domain.ChatRoom, error)
	AddMessage(roomID string, message domain.ChatMessage) error
	GetMessages(roomID string) ([]domain.ChatMessage, error)
}

type chatRoomRepository struct {
	db *redis.Client
}

// NewChatRoomRepository create ChatMessageRepository
func NewChatRoomRepository(client *redis.Client) ChatRoomRepository {
	return chatRoomRepository{client}
}

func (c chatRoomRepository) Create() (*domain.ChatRoom, error) {
	// TODO: multiple room
	room := domain.ChatRoom{ID: "example_room_id"}
	if err := c.db.Set(room.ID, room, 0).Err(); err != nil {
		return nil, err
	}
	return &room, nil
}

func (c chatRoomRepository) Find(roomID string) (*domain.ChatRoom, error) {
	raw, err := c.db.Get(roomID).Result()
	if err != nil {
		return nil, err
	}
	room := domain.ChatRoom{}
	if err := room.UnmarshalBinary(bytes.NewBufferString(raw).Bytes()); err != nil {
		return nil, err
	}

	return &room, nil
}

func (c chatRoomRepository) AddMessage(roomID string, message domain.ChatMessage) error {
	raw, err := c.db.Get(roomID).Result()
	if err != nil {
		return err
	}

	room := domain.ChatRoom{}
	if err := room.UnmarshalBinary(bytes.NewBufferString(raw).Bytes()); err != nil {
		return err
	}
	// TODO: duplication check message id
	room.Messages = append(room.Messages, message)
	if err := c.db.Set(roomID, room, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (c chatRoomRepository) GetMessages(roomID string) ([]domain.ChatMessage, error) {
	raw, err := c.db.Get(roomID).Result()
	if err != nil {
		return nil, err
	}

	room := domain.ChatRoom{}
	if err := room.UnmarshalBinary(bytes.NewBufferString(raw).Bytes()); err != nil {
		return nil, err
	}

	return room.Messages, nil
}
