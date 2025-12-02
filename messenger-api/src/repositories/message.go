package repositories

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/entities"
	"time"

	"github.com/google/uuid"
)

const dateFormat = "2006-01-02"

type MessageDBRegistry struct {
	Id uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	TimesSent int32 `json:"times_sent"`
}

func (m *MessageDBRegistry) ToModel() *entities.Message{
	createdAt, _ := time.Parse(dateFormat, m.CreatedAt)
	return &entities.Message{
		Id: m.Id,
		Content: m.Content,
		CreatedAt: createdAt,
		TimesSent: m.TimesSent,
	}
}

type MessageRespository struct {
	dbConnection db.FileHandler
}

func NewMessageRepository(dbConnection db.FileHandler) *MessageRespository {
	return &MessageRespository{
		dbConnection: dbConnection,
	}
}

func (m *MessageRespository) GetById(id uuid.UUID) (*entities.Message ,error) {
	var messages []MessageDBRegistry

	file, err := m.dbConnection.Read()
	if err != nil {
		return nil, fmt.Errorf("unable to read database file for message lookup: %w", err)
	}

	err = json.Unmarshal(file, &messages)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON format in database file: %w", err)
	}

	for _, message := range messages {
		if message.Id == id {
			return message.ToModel(), nil
		}
	}
	return nil, fmt.Errorf("message with id '%s' was not found in the database", id)
}
