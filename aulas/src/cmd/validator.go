package cmd

import (
	"fmt"
	"spanishGab/aula_camada_model/src/handlers"
)

func ValidateCommandType(input []string) (*handlers.CommandType, error) {
	if len(input) < 2 {
		return nil, fmt.Errorf("error - invalid command")
	}
	commandType := handlers.CommandType(input[1])
	if _, ok := handlers.Commands[commandType]; !ok {
		return nil, fmt.Errorf("error - invalid command")
	}
	return &commandType, nil
}
