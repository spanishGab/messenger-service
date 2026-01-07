package handlers

import "messenger-api/src/repositories"

type MessageHandle struct {
	messageRepository repositories.MessageRespository
}

func newMessageHandle(messageRepository repositories.MessageRespository) *MessageHandle {
	return &MessageHandle{
		messageRepository: messageRepository,
	}
}

