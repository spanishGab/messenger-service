package entities

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	TimesSent int8 `json:"times_sent"`
}

type DateRange struct {
	Start time.Time 
	End time.Time
}

type TimesSent struct {
	Value int8
	Operator string
}

func (ts *TimesSent) MathOperation(value int8) bool {
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
