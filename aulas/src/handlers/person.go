package handlers

import (
	"encoding/json"
	"fmt"
	"spanishGab/aula_camada_model/src/models"
	"spanishGab/aula_camada_model/src/repositories"
	"strconv"
)

type PersonHandler struct {
	personRepository repositories.PersonRepository
}

func NewPersonHandler(personRepository repositories.PersonRepository) *PersonHandler {
	return &PersonHandler{
		personRepository: personRepository,
	}
}

func (ph *PersonHandler) GetPersonById(command Command) (string, error) {
	unparsedID, ok := command.Data["id"]
	if !ok {
		return "", fmt.Errorf("id must be provided")
	}

	id, err := models.ParsePersonID(unparsedID)
	if err != nil {
		return "", fmt.Errorf("the given id is not a valid UUID")
	}

	person, err := ph.personRepository.GetById(*id)
	if err != nil {
		return "", fmt.Errorf("error while searching for person")
	}

	response, err := ph.parseResponse(command, person)
	if err != nil {
		return "", fmt.Errorf("internal application error")
	}
	return string(response), nil
}

func (ph *PersonHandler) GetPersons(command Command) (string, error) {
	unparsedLimit, ok := command.Data["limit"]
	if !ok {
		return "", fmt.Errorf("limit must be provided")
	}
	limit, err := strconv.ParseUint(unparsedLimit, 10, 8)
	if err != nil {
		return "", fmt.Errorf("the given limit is not valid")
	}

	unparsedOffset, ok := command.Data["offset"]
	if !ok {
		return "", fmt.Errorf("offset must be provided")
	}
	offset, err := strconv.ParseUint(unparsedOffset, 10, 8)
	if err != nil {
		return "", fmt.Errorf("the given offset is not valid")
	}

	persons, err := ph.personRepository.GetAll(uint8(limit), uint8(offset))

	response, err := ph.parseResponse(command, persons)
	if err != nil {
		return "", fmt.Errorf("internal application error")
	}
	return string(response), nil
}

func (ph *PersonHandler) parseResponse(command Command, object any) ([]byte, error) {
	format := command.Data["format"]
	var response []byte
	var err error

	var serializeToJSON = func(persons any) ([]byte, error) {
		response, err = json.MarshalIndent(persons, "", "  ")
		if err != nil {
			return nil, err
		}
		return response, nil
	}
	switch OutputFormat(format) {
	case JSONFormat:
		response, err = serializeToJSON(object)
	case Unformatted:
		response, err = json.Marshal(object)
	default:
		response, err = serializeToJSON(object)
	}
	if err != nil {
		return nil, fmt.Errorf("internal application error")
	}
	return response, nil
}
