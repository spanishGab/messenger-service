package cmd

import (
	"fmt"
	"maps"
	"slices"
	"spanishGab/aula_camada_model/src/handlers"
)

const (
	IdOption = "--id"
)

type PersonCommand struct {
	personHandler handlers.PersonHandler
}

func NewPersonCommand(personHandler handlers.PersonHandler) *PersonCommand {
	return &PersonCommand{
		personHandler: personHandler,
	}
}

func (pr *PersonCommand) Run(input []string) {
	commandType, err := ValidateCommandType(input)
	if err != nil {
		fmt.Printf("error - %s", err)
		return
	}

	commandData := input[2:]
	parse, handle, err := pr.chooseParserAndHandler(*commandType)
	if err != nil {
		fmt.Printf("error - %s", err)
		return
	}
	command, err := parse(commandData)
	if err != nil {
		fmt.Printf("error - %s", err)
		return
	}
	result, err := handle(*command)
	if err != nil {
		fmt.Printf("error - %s", err)
		return
	}
	fmt.Println(result)
}

func (pr *PersonCommand) chooseParserAndHandler(commandType handlers.CommandType) (CLIParser, handlers.Handler, error) {
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

func (pr *PersonCommand) parseListCommand(commandData []string) (*handlers.Command, error) {
	var parsedCommandData = make(handlers.CommandData)
	paginationData, err := parsePaginationCommandData(commandData)
	if err != nil {
		return nil, err
	}
	formatData, err := parseFormatCommandData(commandData)
	if err != nil {
		return nil, err
	}
	maps.Copy(parsedCommandData, *paginationData)
	if formatData != nil {
		maps.Copy(parsedCommandData, *formatData)
	}
	return &handlers.Command{
		Type: handlers.List,
		Data: parsedCommandData,
	}, nil
}

func (pr *PersonCommand) parseFindCommand(commandData []string) (*handlers.Command, error) {
	if len(commandData) < 2 {
		return nil, fmt.Errorf("you should provide a person id")
	}
	idIndex := slices.Index(commandData, IdOption)
	if idIndex == -1 {
		return nil, fmt.Errorf("you should provide --id argument")
	}

	var parsedCommandData = handlers.CommandData{
		"id": commandData[idIndex+1],
	}

	formatData, err := parseFormatCommandData(commandData)
	if err != nil {
		return nil, err
	}
	if formatData != nil {
		maps.Copy(parsedCommandData, *formatData)
	}
	return &handlers.Command{
		Type: handlers.Find,
		Data: parsedCommandData,
	}, nil
}
