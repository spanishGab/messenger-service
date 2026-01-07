package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	BirthDate time.Time `json:"birth_date"`
}

func ParsePersonID(unparsedID string) (*uuid.UUID, error) {
	id, err := uuid.Parse(unparsedID)
	if err != nil {
		return nil, fmt.Errorf("invalid person ID: %w", err)
	}
	return &id, nil
}
