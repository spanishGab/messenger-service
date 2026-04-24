package cmd

import (
	"fmt"
	"messenger-api/src/handlers"
)

type MessageCommand struct {
	messageHandler handlers.MessageHandler
}

func NewMessageCommand(messageHandler handlers.MessageHandler) *MessageCommand {
	return &MessageCommand{
		messageHandler: messageHandler,
	}
}

func (mr *MessageCommand) Run(input []string) {
	commandType, err := ValidateCommandType(input)
	if err != nil {
		fmt.Printf("error - %s", err)
		return 
	}

	commandData := input[2:]
	parse, handle, err := ms.chooseParserAndHandler(*commandType)
}

func (ms *MessageCommand) chooseParserAndHandler(commandType handlers.CommandType) (CLIParser, handlers.Handler, error) {
	var parse CLIparser
	var handle handlers.Handler

	switch commandType {
	case handlers.List:
			parse = ms.parseListCommand
	}

}