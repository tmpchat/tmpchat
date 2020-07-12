package domain

import (
	"time"
)

type ChatMessage struct {
	ID string
	Value string
	CreatedAt *time.Time
}
