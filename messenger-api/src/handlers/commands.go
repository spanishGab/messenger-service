package handlers

import "messenger-api/src/shared"

type CommandType string

const (
	Create CommandType = "create"
	Update CommandType = "update"
	List CommandType = "list"
	Find CommandType = "find"
	Delete CommandType = "delete"
)

var Commands shared.Set = shared.Set{
	Create: {},
	Update: {},
	List: {},
	Find: {},
	Delete: {},
}

type CommandData map[string]string

type Command struct {
	Type CommandType
	Data CommandData
}