package domain

import (
	"time"
	"encoding/json"
)

type ChatMessage struct {
	ID string
	Value string
	CreatedAt time.Time
}

func(m ChatMessage) MarshalBinary() ([]byte, error) {
  return json.Marshal(m)
}

func(m ChatMessage) UnmarshalBinary(data []byte) error {
  return json.Unmarshal(data, &m)
}
