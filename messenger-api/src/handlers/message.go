package handlers

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/repositories"

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
