package repository

import (
	"bytes"
	"os"
	"time"

	"github.com/go-redis/redis"

	"github.com/tmpchat/tmpchat/message_broker/domain"
)

type ChatRoomRepository interface {
	Create(roomID string) (*domain.ChatRoom, error)
	Find(roomID string) (*domain.ChatRoom, error)
	AddMessage(roomID string, message domain.ChatMessage) error
	GetMessages(roomID string) ([]domain.ChatMessage, error)
}

type chatRoomRepository struct {
	db *redis.Client
}

// NewChatRoomRepository create ChatMessageRepository
func NewChatRoomRepository() ChatRoomRepository {
	return chatRoomRepository{newRedisClient()}
}

// TODO: config の整理、他の repository でも使用する可能性があるので別のファイルに分ける。
func newRedisClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func (c chatRoomRepository) Create(roomID string) (*domain.ChatRoom, error) {
	now := time.Now()
	room := domain.ChatRoom{ID: roomID, CreatedAt: now, UpdateAt: now}
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
	room, err := c.Find(roomID)
	if err != nil {
		return err
	}
	// TODO: duplication check message id
	room.Messages = append(room.Messages, message)
	room.UpdateAt = time.Now()
	if err := c.db.Set(roomID, room, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (c chatRoomRepository) GetMessages(roomID string) ([]domain.ChatMessage, error) {
	room, err := c.Find(roomID)
	if err != nil {
		return nil, err
	}

	return room.Messages, nil
}
