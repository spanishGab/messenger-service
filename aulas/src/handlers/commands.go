package handlers

import "spanishGab/aula_camada_model/src/shared"

type CommandType string

const (
	Create CommandType = "create"
	Update CommandType = "update"
	Delete CommandType = "delete"
	List   CommandType = "list"
	Find   CommandType = "find"
)

var Commands shared.Set = shared.Set{
	Create: {},
	Update: {},
	Delete: {},
	List:   {},
	Find:   {},
}

type CommandData map[string]string

type Command struct {
	Type CommandType
	Data CommandData
}
