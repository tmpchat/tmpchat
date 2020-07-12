package domain

import (
	"encoding/json"
	"time"
)

type ChatRoom struct {
	ID string
	Messages []ChatMessage
	CreatedAt time.Time
	UpdateAt time.Time
}

func (c ChatRoom) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ChatRoom) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
