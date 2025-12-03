package entities

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	TimesSent int32 `json:"times_sent"`
}

type Filters struct {
	Content string
	CreatedAt string 
	TimesSent *int32 
}
