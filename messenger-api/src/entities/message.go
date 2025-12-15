package entities

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	TimesSent int8 `json:"times_sent"`
}

type DateRange struct {
	Start time.Time 
	End time.Time
}

type Filters struct {
	Content string
	DateRange *DateRange 
	TimesSent *int8 
}
