package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const MAX_CONTENT_SIZE = 250
const MAX_TIMES_SENT_COUNT = 10

type Message struct {
	ID uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	TimesSent uint8 `json:"times_sent"`
}

func ValidateContent(content string) error {
	contentTreated := strings.TrimSpace(content)

	if contentTreated == "" {
		return fmt.Errorf("content cannot be empty")
	}

	if len(content) > MAX_CONTENT_SIZE {
		return fmt.Errorf("'content' must have at most %v characters", MAX_CONTENT_SIZE)
	}

	return nil
}

func ValidateTimesSent(timesSent uint8) error {
	if timesSent > MAX_TIMES_SENT_COUNT {
		return fmt.Errorf("'time sent' must be less than or equal to %v", MAX_TIMES_SENT_COUNT)
	}

	return nil
}

func NewMessage(content string, timesSent uint8) (*Message, error) {
	if err := ValidateContent(content); err != nil {
		return nil, err
	}

	if err := ValidateTimesSent(timesSent); err != nil {
		return nil, err
	}

	return &Message{
		ID: uuid.New(),
		Content: content,
		CreatedAt: time.Now(),
		TimesSent: timesSent,
	}, nil
}

type DateRange struct {
	Start time.Time 
	End *time.Time
}

type TimesSent struct {
	Value uint8
	Operator string
}

func (ts *TimesSent) MatchOperation(value uint8) bool {
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
