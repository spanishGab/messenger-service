package handlers

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/entities"
	"messenger-api/src/repositories"
	"messenger-api/src/shared"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type MessageHandle struct {
	messageRepository repositories.MessageRespository
}

func NewMessageHandle(messageRepository repositories.MessageRespository) *MessageHandle {
	return &MessageHandle{
		messageRepository: messageRepository,
	}
}

func (mh *MessageHandle) GetMessageById(command Command) (string, error) {
	unparsedID, ok := command.Data["id"]
	if !ok {
		return "", fmt.Errorf("id must be provided")
	}

	parsedID, err := uuid.Parse(unparsedID)
	if err != nil {
		return "", fmt.Errorf("id is not valid")
	}

	message, err := mh.messageRepository.GetById(parsedID)
	if err != nil {
		return "", fmt.Errorf("error while searching for message")
	}

	response, err := json.MarshalIndent(message, " ", "")
	if err != nil {
		return "", fmt.Errorf("error formatting response")
	}

	return string(response), nil
}

func (mh *MessageHandle) GetMessages(command Command) (string, error) {
	content, _ := command.Data["content"]

	unparsedCreatedAtStart, _ := command.Data["createdAtStart"]
	var createdAtStart time.Time
	if unparsedCreatedAtStart != "" {
		parsed, err := time.Parse(shared.ShortDateFormat, unparsedCreatedAtStart)
		if err != nil {
			return "", fmt.Errorf("erro ao parsear createdAtStart")
		}
		createdAtStart = parsed
	}

	unparsedCreatedAtEnd, _ := command.Data["createdAtEnd"]
	var createdAtEnd *time.Time
	if unparsedCreatedAtEnd != "" {
		parsed, err := time.Parse(shared.ShortDateFormat, unparsedCreatedAtEnd)
		if err != nil {
			return "", fmt.Errorf("erro ao parsear createdAtEnd")
		}
		createdAtEnd = &parsed
	}

	unparsedTimesSentValue, _ := command.Data["timesSentValue"]
	var timesSentValue uint8
	if unparsedTimesSentValue != "" {
		parsed, err := strconv.ParseUint(unparsedTimesSentValue, 10, 8) 
		if err != nil {
			return "", fmt.Errorf("erro")
		}
		parsedUint8 := uint8(parsed)
		timesSentValue = parsedUint8
	}


	rawTimesSentOperator, _ := command.Data["timesSentOperator"]
	var timesSentOperator string
	if rawTimesSentOperator != "" {
		timesSentOperator = rawTimesSentOperator
	}

	dateRange := entities.DateRange{
		Start: createdAtStart,
		End: createdAtEnd,
	}

	var timesSent entities.TimesSent
	if timesSentValue != 0 && timesSentOperator != "" {
		timesSent = entities.TimesSent{
			Value: timesSentValue,
			Operator: timesSentOperator,
		}
	}

	filters := entities.Filters{
		Content: &content,
		DateRange: &dateRange,
		TimesSent: &timesSent,
	}

	messages, err := mh.messageRepository.GetMessages(filters)
	if err != nil {
		return "", fmt.Errorf("error while searching for message")
	}

	response, err := json.MarshalIndent(messages, " ", "")
	if err != nil {
		return "", fmt.Errorf("error formatting response")
	}

	return string(response), nil
}

