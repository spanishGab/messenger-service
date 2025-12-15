package repositories

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/entities"
	"messenger-api/src/shared"
	"strings"
	"time"

	"github.com/google/uuid"
)


type MessageDBRegistry struct {
	Id uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	TimesSent int8 `json:"times_sent"`
}

func (m *MessageDBRegistry) ToModel() *entities.Message{
	createdAt, _ := time.Parse(shared.ShortDateFormat, m.CreatedAt)
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

func matchTimesSend (value int8, filter entities.Filters) bool {
	switch filter.TimesSent.Operator {
	case "=":
		return value == filter.TimesSent.Value
	case "<":
		return value < filter.TimesSent.Value
	case "<=":
		return value <= filter.TimesSent.Value
	case ">":
		return value > filter.TimesSent.Value
	case ">=":
		return value >= filter.TimesSent.Value
	default: 
		return true
	}
}

func (m *MessageRespository) GetMessages(filters entities.Filters) (*[]entities.Message, error) {
	var messages []MessageDBRegistry
	var results []entities.Message

	file, err := m.dbConnection.Read()
	if err != nil {
		return nil, fmt.Errorf("unable to read database file for message lookup: %w", err)
	}

	err = json.Unmarshal(file, &messages)
		if err != nil {
		return nil, fmt.Errorf("invalid JSON format in database file: %w", err)
	}

	filterContent := strings.ToLower(strings.TrimSpace(*filters.Content))
	filterDateRange := filters.DateRange
	for _, message := range messages {
		if filterContent != "" {
			if !strings.Contains(strings.ToLower(message.Content), filterContent) {
				continue
			}
		}

		if filterDateRange != nil {
			messageCreatedAt, _ := time.Parse(shared.ShortDateFormat, message.CreatedAt)
			if messageCreatedAt.Before(filterDateRange.Start) || 
				messageCreatedAt.After(filterDateRange.End) {
					continue
				}
		}

		if filters.TimesSent != nil {
			if !matchTimesSend(message.TimesSent, filters) {
				continue
			}
		}

		results = append(results, *message.ToModel())
	}

	return &results, nil
}

func (m *MessageRespository) DeleteMessage(id uuid.UUID) error {
	var messages []MessageDBRegistry

	file, err := m.dbConnection.Read()
	if err != nil {
		return fmt.Errorf("unable to read database file for message lookup: %w", err)
	}

	err = json.Unmarshal(file, &messages)
	initialCountItemJSON := len(messages)
	if err != nil {
		return fmt.Errorf("invalid JSON format in database file: %w", err)
	}

	for index, message := range messages {
		if message.Id == id {
			messages = append(messages[:index], messages[index + 1:]...)
		}
	}

	if (initialCountItemJSON) == len(messages) {
		return fmt.Errorf("ID not found %s", id)
	}

	newData, err := json.MarshalIndent(messages, "", " ")
	if err != nil {
		return fmt.Errorf("failed to encode updated data: %w", err)
	}

	_, err = m.dbConnection.Write(newData)
	if err != nil {
		return fmt.Errorf("failed to save database: %w", err)
	}

	return nil
}

func (m *MessageRespository) InsertMessage(message entities.Message) error {
	var messages []MessageDBRegistry

	file, err := m.dbConnection.Read()
	if err != nil {
		return fmt.Errorf("unable to read database file for message lookup: %w", err)
	}

	err = json.Unmarshal(file, &messages)
	if err != nil {
		return fmt.Errorf("error decoding JSON %w", err)
	}

	newMessage := MessageDBRegistry{
		Id: message.Id,
		Content: message.Content,
		CreatedAt: message.CreatedAt.Format(shared.ShortDateFormat),
		TimesSent: message.TimesSent,
	}

	messages = append(messages, newMessage)

	newData, err := json.MarshalIndent(messages, " ", " ")
	if err != nil {
		return fmt.Errorf("failed to encode new data: %w", err)
	}

	_, err = m.dbConnection.Write(newData)
	if err != nil {
		return fmt.Errorf("failed to save database: %w", err)
	}

	return nil
}
