package models

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	BirthDate time.Time `json:"birth_date"`
}
