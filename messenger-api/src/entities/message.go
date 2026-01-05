package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const MAX_CONTENT_SIZE = 250
const MAX_TIME_SENT_COUNT = 10

type Message struct {
	ID uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	TimesSent uint8 `json:"times_sent"`
}

func NewMessage(content string, timeSent uint8) (*Message, error) {
	contentTreated := strings.TrimSpace(content)

	if contentTreated == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}

	if len(content) > MAX_CONTENT_SIZE {
		return nil, fmt.Errorf("'content' must have at most %v characters", MAX_CONTENT_SIZE)
	}

	if timeSent > MAX_TIME_SENT_COUNT {
		return nil, fmt.Errorf("'time sent' must have at most %v characters", MAX_TIME_SENT_COUNT)
	}

	return &Message{
		ID: uuid.New(),
		Content: content,
		CreatedAt: time.Now(),
		TimesSent: timeSent,
	}, nil
}

type DateRange struct {
	Start time.Time 
	End time.Time
}

type TimesSent struct {
	Value uint8
	Operator string
}

func (ts *TimesSent) MathOperation(value uint8) bool {
	switch ts.Operator {
	case "=":
		return value == ts.Value
	case "<":
		return value < ts.Value
	case "<=":
		return value <= ts.Value
	case ">":
		return value > ts.Value
	case ">=":
		return value >= ts.Value
	default: 
		return true
	}
}

type Filters struct {
	Content *string
	DateRange *DateRange 
	TimesSent *TimesSent
}
