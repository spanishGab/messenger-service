package repositories

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/entities"
	"messenger-api/src/shared"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type MessageDBRegistry struct {
	ID uuid.UUID `json:"id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	TimesSent uint8 `json:"times_sent"`
}

type MessageFiltersDTO struct {
	Content *string
	DateRange *entities.DateRange 
	TimesSent *entities.TimesSent
}

type MessageUpdateDTO struct {
	Content *string
	TimesSent *uint8
}

func (m *MessageDBRegistry) ToModel() *entities.Message{
	createdAt, _ := time.Parse(shared.ShortDateFormat, m.CreatedAt)
	return &entities.Message{
		ID: m.ID,
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

func (m *MessageRespository) readTable() ([]MessageDBRegistry, error) {
	var messages []MessageDBRegistry
	file, err := m.dbConnection.Read()

	if err != nil {
		return nil, fmt.Errorf("readTable: unable to read database file for message lookup: %w", err)
	}

	err = json.Unmarshal(file, &messages)
	if err != nil {
		return nil, fmt.Errorf("readTable: invalid JSON format in database file: %w", err)
	}

	return messages, nil
}

func (m *MessageRespository) GetById(id uuid.UUID) (*entities.Message, error) {
  messages, err := m.readTable()
	if err != nil {
		return nil, fmt.Errorf("getById: %w", err)
	}

	for _, message := range messages {
		if message.ID == id {
			return message.ToModel(), nil
		}
	}

	return nil, fmt.Errorf("getById: message with id '%s' was not found in the database", id)
}

func (m *MessageRespository) GetMessages(filters MessageFiltersDTO, pagination Pagination) (*PaginatedResult[entities.Message], error) {
	var data []entities.Message

	messages, err := m.readTable()
	if err != nil {
		return nil, fmt.Errorf("getMessages: %w", err)
	}

	var filterContent string
	if filters.Content != nil {
		filterContent = *filters.Content
	}

	filterDateRange := filters.DateRange
	
	for _, message := range messages {
		if filterContent != "" {
			if !strings.Contains(strings.ToLower(message.Content), filterContent) {
				continue
			}
		}

		if filterDateRange != nil {
			messageCreatedAt, _ := time.Parse(shared.ShortDateFormat, message.CreatedAt)

			if messageCreatedAt.Before(filterDateRange.Start) {
				continue
			}

			if filterDateRange.End != nil && messageCreatedAt.After(*filterDateRange.End) {
				continue
			}
		}

		if filters.TimesSent != nil {
			if !filters.TimesSent.MatchOperation(message.TimesSent) {
				continue
			}
		}

		data = append(data, *message.ToModel())
	}

	results, err := PaginateInMemory(data, pagination)
	if err != nil {
		fmt.Println("getMessages: %w", err)
	}

	return results, nil
}

func (m *MessageRespository) DeleteMessage(id uuid.UUID) error {
	messages, err := m.readTable()
	if err != nil {
		return fmt.Errorf("deleteMessage: %w", err)
	}

	currentMessagesCount := len(messages)

	for index, message := range messages {
		if message.ID == id {
			messages = slices.Delete(messages, index, index + 1)
		}
	}

	if (currentMessagesCount) == len(messages) {
		return fmt.Errorf("deleteMessage: ID not found %s", id)
	}

	newData, err := json.MarshalIndent(messages, "", " ")
	if err != nil {
		return fmt.Errorf("deleteMessage: failed to encode updated data: %w", err)
	}

	_, err = m.dbConnection.Write(newData)
	if err != nil {
		return fmt.Errorf("deleteMessage: failed to save database: %w", err)
	}

	return nil
}

func (m *MessageRespository) InsertMessage(message *entities.Message) error {
	messages, err := m.readTable()
	if err != nil {
		return fmt.Errorf("insertMessage: %w", err)
	}

	newMessage := MessageDBRegistry{
		ID: message.ID,
		Content: message.Content,
		CreatedAt: message.CreatedAt.Format(shared.ShortDateFormat),
		TimesSent: message.TimesSent,
	}

	messages = append(messages, newMessage)

	newData, err := json.MarshalIndent(messages, " ", " ")
	if err != nil {
		return fmt.Errorf("insertMessage: failed to encode new data: %w", err)
	}

	_, err = m.dbConnection.Write(newData)
	if err != nil {
		return fmt.Errorf("insertMessage: failed to save database: %w", err)
	}

	return nil
}

func (m *MessageRespository) UpdateMessage(id uuid.UUID, data *MessageUpdateDTO) error {
	messages, err := m.readTable()
	if err != nil {
		return fmt.Errorf("updateMessage: %w", err)
	}

	for index, message := range messages {
		if id != message.ID { continue }
	
		message.Content = *data.Content
		message.TimesSent = *data.TimesSent
		messages[index] = message
	}

	newData, err := json.MarshalIndent(messages, "", " ")
	if err != nil {
		return fmt.Errorf("updateMessage: marshal %w", err)
	}

	_, err = m.dbConnection.Write(newData)
	if err != nil {
		return fmt.Errorf("updateMessage: failed to save database: %w", err)
	}

	return nil
}
