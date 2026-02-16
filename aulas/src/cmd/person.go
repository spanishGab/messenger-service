package cmd

import (
	"fmt"
	"spanishGab/aula_camada_model/src/handlers"
)

type PersonRouter struct {
	personHandler handlers.PersonHandler
}

func NewPersonRouter(personHandler handlers.PersonHandler) *PersonRouter {
	return &PersonRouter{
		personHandler: personHandler,
	}
}

func (pr *PersonRouter) Route(input []string) {
	commandType, err := ValidateCommandType(input)
	if err != nil {
		fmt.Printf("error - %s", err)
		return
	}

	commandData := input[2:]
	parse, handle, err := pr.ChooseParserAndHandler(*commandType, commandData)
	if err != nil {
		fmt.Printf("error - %s", err)
		return
	}
	command, err := parse(commandData)
	if err != nil {
		fmt.Printf("error - %s", err)
	}
	result, err := handle(*command)
	if err != nil {
		fmt.Printf("error - %s", err)
	}
	fmt.Println(result)
}

func (pr *PersonRouter) ChooseParserAndHandler(commandType handlers.CommandType, commandData []string) (CLIParser, handlers.Handler, error) {
	var parse CLIParser
	var handle handlers.Handler
	switch commandType {
	case handlers.List:
		parse = pr.parseListCommand
		handle = pr.personHandler.GetPersons
	case handlers.Find:
		parse = pr.parseFindCommand
		handle = pr.personHandler.GetPersonById
	default:
		return nil, nil, fmt.Errorf("invalid command")
	}
	return parse, handle, nil
}

func (pr *PersonRouter) parseListCommand(commandData []string) (*handlers.Command, error) {
	if len(commandData) < 2 {
		return nil, fmt.Errorf("you should provide limit and offset arguments")
	}
	return &handlers.Command{
		Type: handlers.List,
		Data: handlers.CommandData{
			"limit":  commandData[0],
			"offset": commandData[1],
		},
	}, nil
}

func (pr *PersonRouter) parseFindCommand(commandData []string) (*handlers.Command, error) {
	if len(commandData) < 1 {
		return nil, fmt.Errorf("you should provide a person id")
	}
	return &handlers.Command{
		Type: handlers.List,
		Data: handlers.CommandData{
			"id": commandData[0],
		},
	}, nil
}
