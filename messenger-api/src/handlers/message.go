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

func (mh *MessageHandle) GetMessageById(command Command) (entities.Result) {
	unparsedID, ok := command.Data["id"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("id must be provided"),
		}
	}

	parsedID, err := uuid.Parse(unparsedID)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("id is not valid"),
		}
	}

	message, err := mh.messageRepository.GetById(parsedID)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error while searching for message"),
		}
	}

	unparsedResponse, err := json.MarshalIndent(message, " ", "")
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error formatting response"),
		}
	}

	response := string(unparsedResponse)
	return entities.Result{
		Value: &response,
	}
}

func (mh *MessageHandle) GetMessages(command Command) (entities.Result) {
	content, _ := command.Data["content"]

	unparsedCreatedAtStart, _ := command.Data["createdAtStart"]
	var createdAtStart time.Time
	if unparsedCreatedAtStart != "" {
		parsed, err := time.Parse(shared.ShortDateFormat, unparsedCreatedAtStart)
		if err != nil {
			return entities.Result{
				Error: fmt.Errorf("error parsing createdAtStart"),
			}
		}
		createdAtStart = parsed
	}

	unparsedCreatedAtEnd, _ := command.Data["createdAtEnd"]
	var createdAtEnd *time.Time
	if unparsedCreatedAtEnd != "" {
		parsed, err := time.Parse(shared.ShortDateFormat, unparsedCreatedAtEnd)
		if err != nil {
			return entities.Result{
				Error: fmt.Errorf("error parsing createdAtEnd"),
			}
		}
		createdAtEnd = &parsed
	}

	unparsedTimesSentValue, _ := command.Data["timesSentValue"]
	var timesSentValue uint8
	if unparsedTimesSentValue != "" {
		parsed, err := strconv.ParseUint(unparsedTimesSentValue, 10, 8) 
		if err != nil {
			return entities.Result{
				Error: fmt.Errorf("error parsing timesSentValue"),
			}
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
		return entities.Result{
			Error: fmt.Errorf("error while searching for message"),
		}
	}

	unparsedResponse, err := json.MarshalIndent(messages, " ", "")
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error formatting response"),
		}
	}

	response := string(unparsedResponse)
	return entities.Result{
		Value: &response,
	}
}

func (mh *MessageHandle) DeleteMessage(command Command) (entities.Result) {
	unparsedId, ok := command.Data["id"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("id must be provided"),
		}
	}

	id, err := uuid.Parse(unparsedId)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error parsing id"),
		}
	}

	err = mh.messageRepository.DeleteMessage(id)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error deleting message"),
		}
	}

	unparsedResponse := fmt.Sprintf("Message from ID %s deleted successfully", id)
	response := string(unparsedResponse)
	return entities.Result{
		Value: &response,
	}
}

func (mh *MessageHandle) InsertMessage(command Command) (entities.Result) {
	content, ok := command.Data["content"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("content must be provided"),
		}
	}

	unparsedTimesSent, ok := command.Data["timesSent"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("timesSent must be provided"),
		}
	}

	parsedTimesSent, err := strconv.ParseUint(unparsedTimesSent, 10, 8)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error parsing timesSent"),
		}
	}

	timesSent := uint8(parsedTimesSent)
	messange, err := entities.NewMessage(content, timesSent)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error inserting message"),
		}
	}

	err = mh.messageRepository.InsertMessage(messange)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error insert message"),
		}
	}

	unparsedResponse := fmt.Sprintf("Message from created successfully")
	response := string(unparsedResponse)
	return entities.Result{
		Value: &response,
	}
}

func (mh *MessageHandle) UpdateMessage(command Command) (entities.Result) {
	unparsedID, ok := command.Data["id"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("id must be provided"),
		}
	}

	id, err := uuid.Parse(unparsedID)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("id is not valid"),
		}
	}

	content, ok := command.Data["content"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("content must be provided"),
		}
	}

	unparsedTimesSent, ok := command.Data["timesSent"]
	if !ok {
		return entities.Result{
			Error: fmt.Errorf("timesSent must be provided"),
		}
	}

	parsedTimesSent, err := strconv.ParseUint(unparsedTimesSent, 10, 8) 
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error parsing timesSent"),
		}
	}

	timesSent := uint8(parsedTimesSent)

	data := &entities.MessageUpdate{
		Content: &content,
		TimesSent: &timesSent,
	}

	err = mh.messageRepository.UpdateMessage(id, data)
	if err != nil {
		return entities.Result{
			Error: fmt.Errorf("error updated message"),
		}
	}

	unparsedResponse := fmt.Sprintf("Message from ID %s updated successfully", id)
	response := string(unparsedResponse)
	return entities.Result{
		Value: &response,
	}
}
